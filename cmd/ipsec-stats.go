// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package ipseccmd implements the CLI commands
package cmd

import (
	"fmt"

	"github.com/opiproject/godpu/pkg/ipsec"
	"github.com/spf13/cobra"
)

func NewStatsCommand() *cobra.Command {
	var (
		addr string
	)
	cmd := &cobra.Command{
		Use:     "stats",
		Aliases: []string{"c"},
		Short:   "Queries ipsec statistics",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			res := ipsec.Stats(addr)
			fmt.Println(res)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&addr, "addr", "localhost:50151", "address or OPI gRPC server")
	return cmd
}
