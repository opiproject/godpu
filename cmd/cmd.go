// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package cmd implements the ipsec related CLI commands
package cmd

import (
	"log"
	"os"

	"github.com/opiproject/godpu/cmd/common"
	"github.com/opiproject/godpu/cmd/evpnipsec"
	"github.com/opiproject/godpu/cmd/inventory"
	"github.com/opiproject/godpu/cmd/ipsec"
	"github.com/opiproject/godpu/cmd/network"
	"github.com/opiproject/godpu/cmd/storage"
	"github.com/spf13/cobra"
)

// NewCommand handles the cli for evpn, ipsec, invetory and storage
func NewCommand() *cobra.Command {
	//
	// This is the root command for the CLI
	//

	c := &cobra.Command{
		Use:   "godpu",
		Short: "godpu - DPUs and IPUs cli commands",
		Run: func(cmd *cobra.Command, _ []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatalf("[ERROR] %s", err.Error())
			}
			os.Exit(1)
		},
	}
	c.AddCommand(inventory.NewInventoryCommand())
	c.AddCommand(ipsec.NewIPSecCommand())
	c.AddCommand(storage.NewStorageCommand())
	c.AddCommand(network.NewEvpnCommand())
	c.AddCommand(evpnipsec.NewEvpnIPSecCommand())
	flags := c.PersistentFlags()
	flags.String(common.AddrCmdLineArg, "localhost:50151", "address of OPI gRPC server")
	flags.String(common.TLSFiles, "", "TLS files in client_cert:client_key:ca_cert format.")

	return c
}
