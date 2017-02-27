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

import (
	"fmt"

	"github.com/spf13/cobra"
)

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "Perform operations on API routes",
	Long: `Routes are API endpoint exposed by
Limitless and the loaded modules.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to select an operation. See -h for help.")
	},
}

func init() {
	RootCmd.AddCommand(routeCmd)
}
