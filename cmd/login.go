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

	"github.com/spf13/cobra"
)

var username string
var password string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log a user in to a Project Limitless installation",
	Long: `Log a user in to a Project Limitless installation
as specified in the configuration file.

If login succeeds the access token is stored and used for all calls, including
calls that require authentication.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("login called for username: " + username)
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().StringVar(&username, "username", "", "The user's username")
	loginCmd.PersistentFlags().StringVar(&password, "password", "", "The user's password")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
