// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package ipseccmd implements the CLI commands
package ipseccmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewIpsecCommand returns the ipsec command
func NewIpsecCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ipsec",
		Short: "DPUs and IPUs ipsec commands",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	cmd.AddCommand(newStatsCommand())
	cmd.AddCommand(newTestCommand())
	// cmd.AddCommand(newBridgeRm())

	return cmd
}
