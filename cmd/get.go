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

import "github.com/spf13/cobra"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets API objects",
	Long: `Get returns API objects and displays them as a list.

The command following 'get' specifies the object to fetch.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Errorf("You need to specify an API object to get. See `%s --help`", RootCmd.Use)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
