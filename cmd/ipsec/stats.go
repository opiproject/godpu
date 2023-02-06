// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package ipseccmd implements the CLI commands
package ipseccmd

import (
	"fmt"

	"github.com/opiproject/godpu/pkg/ipsec"
	"github.com/spf13/cobra"
)

func newStatsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "stats",
		Aliases: []string{"c"},
		Short:   "Queries ipsec statistics",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			res := ipsec.Stats()
			fmt.Println(res)
		},
	}
	return cmd
}
