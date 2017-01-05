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

var wide bool

// routesCmd represents the routes command
var routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "Returns a list of available API routes",
	Long: `Returns a list of available API routes along with the path, description,
HTTP method and a boolean specifying if authentication is required for the path`,
	Run: func(cmd *cobra.Command, args []string) {
		apiRoutes, response, err := api.AdminRoutesGet()
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
		writeFields(tw, true, "Method", "Path", "Description", "Authenticated")
		for _, route := range apiRoutes {
			writeFields(tw, false, route.Method, route.Path, route.Description, fmt.Sprintf("%v", route.RequiresAuthentication))
		}
		tw.Flush()
	},
}

func init() {
	getCmd.AddCommand(routesCmd)
}
