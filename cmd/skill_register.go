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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var skillFilename string

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new skill",
	Long: `Registers a new skill by sending the
Skill JSON registration info to the Limitless API.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		if skillFilename == "" {
			logger.Error("You must provide a path to a Skill JSON file")
			return
		}
		if _, err := os.Stat(skillFilename); os.IsNotExist(err) {
			logger.Errorf("Unable to open Skill JSON file '%s'", skillFilename)
			return
		}

		fileBytes, err := ioutil.ReadFile(skillFilename)
		if err != nil {
			logger.Error("Unable to read Skill JSON file '%s': %s", skillFilename, err.Error())
			return
		}

		_ = fileBytes
		// TODO: Unmarshal into Skill?

	},
}

func init() {
	skillCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringVarP(&skillFilename, "filename", "f", "", "The Skill JSON file to register")
}
