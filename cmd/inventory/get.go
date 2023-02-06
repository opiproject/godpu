// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package inventory implements the CLI commands
package inventory

import (
	"fmt"
	"os"

	"github.com/opiproject/godpu/pkg/inventory"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Gets DPU inventory information",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := inventory.Get()
		fmt.Println(res)
	},
}

// Initialize executes the get command.
func Initialize() {
	if err := getCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
