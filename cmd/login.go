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
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ProjectLimitless/llctl/swagger"
	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
)

var username string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log a user in to a Project Limitless installation",
	Long: `Log a user in to a Project Limitless installation
as specified in the configuration file.

If login succeeds the access token is stored and used for all calls, including
calls that require authentication.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: find a better way to handle required flags
		if username == "" {
			logger.Error("A username is required for login")
			return
		}

		fmt.Printf("Password: ")
		// Windows users might not be used to not see any feedback when
		// entering a password like 'GetPasswd'. Rather using *** masked input.
		passwordBytes, err := gopass.GetPasswdMasked()
		if err != nil {
			logger.Errorf("Unable to read password: %s", err.Error())
			return
		}
		password := string(passwordBytes)

		loginResponse, response, err := api.LoginPost(swagger.LoginCredentials{
			Username: username,
			Password: password,
		})
		if err != nil {
			logger.Errorf("Unable to call API: %s", err.Error())
		}

		if isFailedStatus(response.StatusCode) {
			handleErrorResponse(response)
			logger.Critical("Login Failed. Incorrect username or password.")
			return
		}

		if isDebug {
			fmt.Println(prettyJSON(response.Payload))
		}

		// Save the loginResult in .llcache
		llcache, err := json.Marshal(&loginResponse)
		if err != nil {
			logger.Critical("Unable to marshal LoginResponse: %s", err.Error())
			return
		}
		err = ioutil.WriteFile(cacheFile, llcache, 0644)
		if err != nil {
			logger.Critical("Unable to write login result to '%s': %s", cacheFile, err.Error())
			return
		}

		logger.Info("Login successful. You may now call authenticated routes.")
		logger.Infof("Authenticated as %s %s (%s)",
			loginResponse.Name,
			loginResponse.Surname,
			loginResponse.UserName,
		)
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "The user's username")
	loginCmd.MarkFlagRequired("username")

}
