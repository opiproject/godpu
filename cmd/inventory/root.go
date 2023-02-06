// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package inventorycmd implements the CLI commands
package inventorycmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewInventoryCommand returns the inventory command
func NewInventoryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inventory",
		Short: "DPUs and IPUs inventory commands",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	cmd.AddCommand(newGetCommand())
	// cmd.AddCommand(newBridgeRm())
	// cmd.AddCommand(newBridgeRm())

	return cmd
}
