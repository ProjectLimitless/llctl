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
			fmt.Println("The username flag is required for login")
			return
		}

		fmt.Printf("Password: ")
		// Windows users might not be used to not see any feedback when
		// entering a password like 'GetPasswd'. Rather using *** masked input.
		passwordBytes, err := gopass.GetPasswdMasked()
		if err != nil {
			fmt.Println("Unable to read password:", err.Error())
			return
		}
		password := string(passwordBytes)

		_ = password

		api := swagger.NewDefaultApiWithBasePath("http://127.0.0.1:8080")
		response, err := api.ApiLoginPost(swagger.LoginCredentials{
			Username: username,
			Password: password,
		})

		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(response.Payload))

		// TODO: Add responses to swagger definition
		// to read more data from the api
		// on login, save the token in .llcache?
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "The user's username")
	loginCmd.MarkFlagRequired("username")

}
