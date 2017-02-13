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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// skillCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// skillCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
