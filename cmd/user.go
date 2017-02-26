// This file is part of llctl.
// Copyright © 2017 Donovan Solms.
// Project Limitless
// https://www.projectlimitless.io
//
// llctl and Project Limitless is free software: you can redistribute it and/or modify
// it under the terms of the Apache License Version 2.0.
//
// You should have received a copy of the Apache License Version 2.0 with
// llctl. If not, see http://www.apache.org/licenses/LICENSE-2.0.

package cmd

import "github.com/spf13/cobra"

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Perform operations on users",
	Long: `Users are allowed to access and
interact with the Limitless installation`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(userCmd)
}
