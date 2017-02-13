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

var skillUUID string

// deregisterCmd represents the deregister command
var deregisterCmd = &cobra.Command{
	Use:   "deregister",
	Short: "Deregister a registered skill",
	Long: `Deregisters a previously registered skill by sending the
Skill UUID to the Limitless API.`,
	Run: func(cmd *cobra.Command, args []string) {
		if skillUUID == "" {
			logger.Error("You must provide a valid Skill UUID")
			return
		}
		response, err := api.SkillsIdDelete(skillUUID)
		if err != nil {
			handleErrorResponse(response)
			return
		}
		logger.Infof("The skill with UUID '%s' was deregistered", skillUUID)
	},
}

func init() {
	skillCmd.AddCommand(deregisterCmd)
	deregisterCmd.Flags().StringVarP(&skillUUID, "uuid", "u", "", "The Skill UUID to deregister")
}
