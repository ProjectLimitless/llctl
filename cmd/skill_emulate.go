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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ProjectLimitless/llctl/swagger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tylerb/graceful"
)

// emulateCmd represents the emulate command that registers
// a skill and emulates the skill's endpoint.
var emulateCmd = &cobra.Command{
	Use:   "emulate",
	Short: "Register and emulate a skill endpoint once",
	Long: `Register the given skill and set the endpoint
to the given IP and port.

Executes a server and waits for skill a single call from Project
Limitless, then shuts down.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Lookup("host").Changed == false &&
			viper.GetString("skills.emulateHost") != "" {
			host = viper.GetString("skills.emulateHost")
		}
		logger.Infof("The Skill emulation server is running on '%s'", host)

		// Register skill with the 'host' as the endpoint
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
			logger.Error("Unable to read Skill JSON file '%s': %s",
				skillFilename, err.Error())
			return
		}

		var skill swagger.Skill
		err = json.Unmarshal(fileBytes, &skill)
		if err != nil {
			logger.Errorf("The given file does not contain a valid skill: %s",
				err.Error())
			return
		}
		// Ensure we are using the network binding,
		// even if it is set in the file
		skill.Binding = "Network"
		skill.Executor.Url = "http://" + host
		skill.Executor.ValidateCertificate = false

		registrationDetails, response, err := api.SkillsPost(skill)
		if err != nil {
			handleErrorResponse(response)
			return
		}
		if registrationDetails.Registered == false {
			logger.Error("Unable to register skill, see response dump below")
			fmt.Println(prettyJSON(response.Payload))
			return
		}
		logger.Infof("Skill '%s' has been registered", registrationDetails.UUID)

		// Run the server and shutdown after any input is received
		server := &graceful.Server{
			Timeout: time.Second * 10,
			Server: &http.Server{
				Addr: host,
			},
		}
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			contentType := req.Header.Get("Content-Type")
			if strings.Contains(contentType, "application/json") == false {
				logger.Errorf("Unable to handle incoming Content-Type '%s'",
					contentType)
				return
			}

			decoder := json.NewDecoder(req.Body)
			var payload swagger.SkillNetworkExecutorPayload
			err := decoder.Decode(&payload)
			if err != nil {
				logger.Errorf("Unable to parse NetworkExecutorPayload: %s", err.Error())
				return
			}
			defer req.Body.Close()

			logger.Infof("Received Skill Executor Call for skill '%s'", payload.SkillUUID)

			response := "Yes"
			logger.Infof("Responded with '%s'", response)

			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Language", "en-ZA")
			w.Write([]byte(response))

			server.Stop(time.Second)
		})
		server.Handler = mux
		err = server.ListenAndServe()
		if err != nil {
			logger.Errorf("Unable to start emulation server: %s", err.Error())
		}
	},
}

func init() {
	skillCmd.AddCommand(emulateCmd)
	emulateCmd.Flags().StringVarP(&skillFilename, "filename", "f", "",
		"The Skill JSON file to register and emulate")
	emulateCmd.Flags().StringVarP(&host, "host", "H", "0.0.0.0:9000",
		"The Skill endpoint address")
}
