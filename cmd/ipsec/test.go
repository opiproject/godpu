// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package ipseccmd implements the CLI commands
package ipseccmd

import (
	"fmt"

	"github.com/opiproject/godpu/pkg/ipsec"
	"github.com/spf13/cobra"
)

func newTestCommand() *cobra.Command {
	var (
		addr     string
		pingaddr string
	)
	cmd := &cobra.Command{
		Use:     "test",
		Aliases: []string{"c"},
		Short:   "Test ipsec functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			res := ipsec.TestIpsec(addr, pingaddr)
			fmt.Println(res)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&addr, "addr", "localhost:50151", "address or OPI gRPC server")
	flags.StringVar(&pingaddr, "pingaddr", "localhost", "address of other tunnel end to Ping")
	return cmd
}
