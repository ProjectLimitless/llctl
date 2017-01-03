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

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Tests the connection to a Project Limitless instance",
	Long: `Ping creates a connection to a Project Limitless instance
and attempts to authenticate to it.

If connection and authentication succeeds, 'Pong!' is returned, otherwise
the error will be displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("ping called")
		fmt.Println(host)
		fmt.Println(port)
	},
}

func init() {
	RootCmd.AddCommand(pingCmd)
}
