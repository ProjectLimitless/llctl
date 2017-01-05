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
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// modulesCmd represents the modules command
var modulesCmd = &cobra.Command{
	Use:   "modules",
	Short: "Returns a list of loaded modules",
	Long: `Returns a list of loaded moodules along with the description
and version`,
	Run: func(cmd *cobra.Command, args []string) {
		modules, response, err := api.AdminModulesGet()
		if err != nil {
			logger.Errorf("Unable to call API: %s", err.Error())
		}
		if isFailedStatus(response.StatusCode) {
			handleErrorResponse(response)
			return
		}

		if isDebug {
			fmt.Println(prettyJSON(response.Payload))
		}

		tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		if wide {
			writeFields(tw, true, "Type", "Title", "Description", "Version", "Author")
			for _, module := range modules {
				writeFields(tw, false, module.Type_, module.Title, module.Description, module.Version, module.Author)
			}
		} else {
			writeFields(tw, true, "Type", "Title", "Version", "Author")
			for _, module := range modules {
				writeFields(tw, false, module.Type_, module.Title, module.Version, module.Author)
			}
		}
		tw.Flush()
	},
}

func init() {
	getCmd.AddCommand(modulesCmd)
	modulesCmd.PersistentFlags().BoolVarP(&wide, "wide", "w", false, "Print more details")
}
