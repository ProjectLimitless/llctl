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

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Returns a list of registered users",
	Long: `Returns a list of registered users on this
Project Limitless installation`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("users called")
	},
}

func init() {
	getCmd.AddCommand(usersCmd)
}
