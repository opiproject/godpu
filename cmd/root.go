// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package commands implements the CLI commands
package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	inventorycmd "github.com/opiproject/godpu/cmd/inventory"
	ipseccmd "github.com/opiproject/godpu/cmd/ipsec"
)

// NewRootCommand returns the root command
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "godpu",
		Short: "godpu - DPUs and IPUs cli commands",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.AddCommand(ipseccmd.NewIpsecCommand())
	cmd.AddCommand(inventorycmd.NewInventoryCommand())

	return cmd
}

// Execute executes the root command.
func Execute() {
	if err := NewRootCommand().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
