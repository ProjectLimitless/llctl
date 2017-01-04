// This file is part of llctl.
// Copyright © 2016 Donovan Solms.
// Project Limitless
// https://www.projectlimitless.io
//
// llctl and Project Limitless is free software: you can redistribute it and/or modify
// it under the terms of the Apache License Version 2.0.
//
// You should have received a copy of the Apache License Version 2.0 with
// llctl. If not, see http://www.apache.org/licenses/LICENSE-2.0.

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ProjectLimitless/llctl/swagger"
	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string
var apiEndpoint string
var api *swagger.DefaultApi

var isDebug bool

var logger *logging.Logger
var cacheFile string
var user swagger.LoginResponse

// buildVersion is the version of the 'llctl'
// tool which is updated by automated builds.
var buildVersion string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "llctl",
	Short: "Project Limitless Control",
	Long: `'llctl' is a command line interface for running commands against a Project
Limitless installation.

See https://docs.projectlimitless.io/llctl for more information.

Project Limitless - Everyone's Jarvis
https://www.projectlimitless.io
`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	buildVersion = version
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "./.llctl.yaml", "config file")
	RootCmd.PersistentFlags().StringVar(&apiEndpoint, "endpoint", "http://127.0.0.1:8762/api", "The Project Limitless builtin API endpoint")
	RootCmd.PersistentFlags().BoolVarP(&isDebug, "debug", "d", false, "Enable debug output")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	cacheFile = "./.llcache"
	// Enable ability to specify config file via flag
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}

	// Name of config file (without extension)
	viper.SetConfigName(".llctl")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	if viper.IsSet("debug") {
		isDebug = viper.GetBool("debug")
	}
	if viper.IsSet("api.endpoint") {
		apiEndpoint = viper.GetString("api.endpoint")
	}

	// Setup logging with log level
	logger = logging.MustGetLogger("llctl")
	format := logging.MustStringFormatter(`%{color}%{time:15:04:05} %{level:.5s} > %{message}%{color:reset}`)
	stdout := logging.NewLogBackend(os.Stdout, "", 0)
	stdoutFormatter := logging.NewBackendFormatter(stdout, format)
	stdoutLogLevel := logging.AddModuleLevel(stdoutFormatter)

	if isDebug {
		stdoutLogLevel.SetLevel(logging.DEBUG, "")
	} else {
		stdoutLogLevel.SetLevel(logging.INFO, "")
	}
	logging.SetBackend(stdoutLogLevel)

	logger.Debug("Debugging is enabled")

	// Setup Swagger generated API code. Spec available at
	// https://app.swaggerhub.com/api/projectlimitless/builtin-api
	api = swagger.NewDefaultApiWithBasePath(apiEndpoint)
	api.Configuration.Debug = isDebug

	// Read cached user data
	llcache, err := ioutil.ReadFile(cacheFile)
	if err != nil {
		logger.Debugf("Limitless cache file '%s' is invalid or doesn't exist", cacheFile)
	} else {
		err = json.Unmarshal(llcache, &user)
		if err != nil {
			_ = os.Remove(cacheFile)
			logger.Debugf("Limitless cache file '%s' is invalid. Removed.", cacheFile)
		} else {
			api.Configuration.APIKey["Authorization"] = user.AccessToken
		}
	}
}

// isFailedStatus checks is a given HTTP status code is
// not a success status. 2XX statuses are considered success.
func isFailedStatus(statusCode int) bool {
	if statusCode >= 200 && statusCode < 300 {
		return false
	}
	return true
}

// handleErrorResponse defines a uniform way of handling
// error responses from the API
func handleErrorResponse(response *swagger.APIResponse) {
	if len(response.Payload) > 2 {
		fmt.Println(prettyJSON(response.Payload))
	}
	logger.Errorf("Received Status '%d' from the API: %s", response.StatusCode, response.Status)
	if response.StatusCode == 401 {
		logger.Warningf("You need to log in to access this route. See '%s %s --help'", RootCmd.Use, "login")
	}
}

// prettyJSON format JSON to be more readable
func prettyJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data, "", "\t")
	if err != nil {
		logger.Criticalf("Unable to parse JSON: %s", err.Error())
		return ""
	}
	return string(prettyJSON.Bytes())
}
