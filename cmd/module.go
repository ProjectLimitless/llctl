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

// moduleCmd represents the modules command
var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Perform operations on modules",
	Long: `Modules are binary loaded libraries
that the Limitless core uses.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(moduleCmd)
}
