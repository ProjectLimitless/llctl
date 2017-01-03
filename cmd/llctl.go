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
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string
var host net.IP
var port uint32

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

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&configFile, "config", "./.llctl.yaml", "config file")
	RootCmd.PersistentFlags().IPVar(&host, "host", net.ParseIP("127.0.0.1"), "The Project Limitless host to connect to")
	RootCmd.PersistentFlags().Uint32Var(&port, "port", 8762, "The Project Limitless port to connect to")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if configFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(configFile)
	}

	viper.SetConfigName(".llctl") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
