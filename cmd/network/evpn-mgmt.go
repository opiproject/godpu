// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (c) 2024 Ericsson AB.

// Package network implements the network related CLI commands
package network

import (
	"context"
	"log"
	"time"

	"github.com/opiproject/godpu/cmd/common"
	"github.com/opiproject/godpu/network"
	"github.com/spf13/cobra"
)

// DumpNetlinkDatabase Get netlink database details
func DumpNetlinkDatabase() *cobra.Command {
	var details bool

	cmd := &cobra.Command{
		Use:   "dump-netlink-DB",
		Short: "Show details of a netlink database",
		Long:  "Show details of  netlink database with current running config",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewManagement(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			result, err := evpnClient.DumpNetlinkDB(ctx, details)
			if err != nil {
				log.Fatalf("DumpNetlinkDatabase: Error occurred while dumping the netlink database: %q", err)
			}
			log.Printf("DumpNetlinkDatabase: %s", result.Details)
		},
	}

	cmd.Flags().BoolVar(&details, "details", false, "get the dump with details")

	if err := cmd.MarkFlagRequired("details"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}
