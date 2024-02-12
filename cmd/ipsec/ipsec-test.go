// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package ipsec implements the ipsec related CLI commands
package ipsec

import (
	"fmt"

	"github.com/opiproject/godpu/ipsec"
	"github.com/spf13/cobra"
)

// NewTestCommand returns the ipsec tests command
func NewTestCommand() *cobra.Command {
	var (
		addr     string
		pingaddr string
	)
	cmd := &cobra.Command{
		Use:     "test",
		Aliases: []string{"c"},
		Short:   "Test ipsec functionality",
		Args:    cobra.NoArgs,
		Run: func(_ *cobra.Command, _ []string) {
			res := ipsec.TestIpsec(addr, pingaddr)
			fmt.Println(res)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&addr, "addr", "localhost:50151", "address or OPI gRPC server")
	flags.StringVar(&pingaddr, "pingaddr", "localhost", "address of other tunnel end to Ping")
	return cmd
}
