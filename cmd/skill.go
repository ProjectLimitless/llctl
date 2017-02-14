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

// The host for the skill server
var host string

// The filename of the skill JSON file to register
var skillFilename string

// skillCmd represents the skill command
var skillCmd = &cobra.Command{
	Use:   "skill",
	Short: "Perform operations on a skill",
	Long: `Skills are new functionality that can
be added to Project Limitless.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
	},
}

func init() {
	RootCmd.AddCommand(skillCmd)
}
