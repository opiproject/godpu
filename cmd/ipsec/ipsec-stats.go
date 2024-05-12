// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package ipsec implements the ipsec related CLI commands
package ipsec

import (
	"fmt"
	"log"

	"github.com/opiproject/godpu/cmd/common"
	"github.com/opiproject/godpu/ipsec"
	"github.com/spf13/cobra"
)

// NewStatsCommand returns the ipsec stats command
func NewStatsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "stats",
		Aliases: []string{"c"},
		Short:   "Queries ipsec statistics",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			res := ipsec.Stats(addr)
			fmt.Println(res)
		},
	}
	return cmd
}

// NewIPSecCommand tests the  inventory
func NewIPSecCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ipsec",
		Aliases: []string{"g"},
		Short:   "Tests ipsec functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatalf("[ERROR] %s", err.Error())
			}
		},
	}

	cmd.AddCommand(NewStatsCommand())
	cmd.AddCommand(NewTestCommand())
	return cmd
}
