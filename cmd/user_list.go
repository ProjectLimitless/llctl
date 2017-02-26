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
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// userListCmd represents the list command
var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "Returns a list of registered users",
	Long: `Returns a list of registered users for this
Project Limitless installation`,
	Run: func(cmd *cobra.Command, args []string) {
		users, response, err := api.AdminUsersGet()
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
			writeFields(tw, true, "ID", "Username", "First name", "Last name", "Created", "Deleted")
			for _, user := range users {
				writeFields(tw, false, fmt.Sprintf("%d", user.ID), user.Username, user.FirstName, user.LastName, user.DateCreated.Format("2006-01-02 15:04:05"), fmt.Sprintf("%v", user.IsDeleted))
			}
		} else {
			writeFields(tw, true, "ID", "Username", "First name", "Last name")
			for _, user := range users {
				writeFields(tw, false, fmt.Sprintf("%d", user.ID), user.Username, user.FirstName, user.LastName)
			}
		}
		tw.Flush()
	},
}

func init() {
	userListCmd.PersistentFlags().BoolVarP(&wide, "wide", "w", false, "Print more details")
	userCmd.AddCommand(userListCmd)
}
