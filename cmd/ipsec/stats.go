// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package ipsec implements the CLI commands
package ipsec

import (
	"fmt"
	"os"

	"github.com/opiproject/godpu/pkg/ipsec"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:     "stats",
	Aliases: []string{"c"},
	Short:   "Queries ipsec statistics",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := ipsec.Stats()
		fmt.Println(res)
	},
}

// Initialize executes the stats command.
func Initialize() {
	if err := statsCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
