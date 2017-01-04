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

// modulesCmd represents the modules command
var modulesCmd = &cobra.Command{
	Use:   "modules",
	Short: "Returns a list of loaded modules",
	Long: `Returns a list of loaded moodules along with the description
and version`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("modules called")
	},
}

func init() {
	getCmd.AddCommand(modulesCmd)
}
