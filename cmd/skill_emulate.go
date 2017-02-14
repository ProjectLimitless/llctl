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
	"github.com/spf13/viper"
)

// emulateCmd represents the emulate command that registers
// a skill and emulates the skill's endpoint.
var emulateCmd = &cobra.Command{
	Use:   "emulate",
	Short: "Register and emulate a skill endpoint",
	Long: `Register the given skill and set the endpoint
to the given IP and port.

Executes a server and waits for skill calls from Project
Limitless.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("emulate called")

		// TODO: Check if host is set
		host = viper.GetString("skills.emulateHost")
		logger.Debugf("Skill emulate host set as '%s'", host)
	},
}

func init() {
	skillCmd.AddCommand(emulateCmd)
	emulateCmd.Flags().StringVarP(&skillFilename, "filename", "f", "",
		"The Skill JSON file to register and emulate")
	emulateCmd.Flags().StringVarP(&host, "host", "H", "http://0.0.0.0:9000",
		"The Skill endpoint address")

}
