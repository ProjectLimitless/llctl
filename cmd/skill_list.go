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

// listCmd represents the list of skills command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the registered skills",
	Long: `Lists the currently registered skills
for the given Limitless installation`,
	Run: func(cmd *cobra.Command, args []string) {
		skills, response, err := api.SkillsGet()
		if err != nil {
			handleErrorResponse(response)
			return
		}
		plural := "s"
		if len(skills) == 1 {
			plural = ""
		}
		logger.Infof("%d skill%s registered", len(skills), plural)

		for i, skill := range skills {
			logger.Infof("Skill #%d - '%s (%s)' by '%s': %s", i+1, skill.Name,
				skill.UUID, skill.Author, skill.ShortDescription)
		}
	},
}

func init() {
	skillCmd.AddCommand(listCmd)
}
