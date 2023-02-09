// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package main implements the CLI commands
package main

import (
	cmd "github.com/opiproject/godpu/cmd"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	command := newCommand()
	if err := command.Execute(); err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
		os.Exit(-1)
	}
}

func newCommand() *cobra.Command {
	//
	// This is the root command for the CLI
	//

	c := &cobra.Command{
		Use:   "godpu",
		Short: "godpu - DPUs and IPUs cli commands",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatalf("[ERROR] %s", err.Error())
				os.Exit(1)
			}
			os.Exit(1)
		},
	}
	c.AddCommand(cmd.NewGetCommand())
	c.AddCommand(cmd.NewTestCommand())
	c.AddCommand(cmd.NewStatsCommand())
	return c
}
