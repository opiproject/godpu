// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package inventorycmd implements the CLI commands
package cmd

import (
	"fmt"

	"github.com/opiproject/godpu/pkg/inventory"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "Gets DPU inventory information",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			res := inventory.Get()
			fmt.Println(res)
		},
	}
	return cmd
}
