// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2025 Dell Inc, or its subsidiaries.

// Package netintf implements the network interface related CLI commands
package netintf

import (
	"context"
	"log"

	"github.com/opiproject/godpu/cmd/common"
	network "github.com/opiproject/godpu/network"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

// ListNetworkInterfaces lists all Network Interface details from OPI server
func ListNetworkInterfaces() *cobra.Command {
	var pageSize int32
	var pageToken string

	cmd := &cobra.Command{
		Use:   "list-net-interfaces",
		Short: "List the network interfaces",
		Run: func(c *cobra.Command, _ []string) {
			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(common.TimeoutCmdLineArg)
			cobra.CheckErr(err)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)

			netifClient, err := network.NewNetInterface(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			for {
				resp, err := netifClient.ListNetInterfaces(ctx, pageSize, pageToken)
				if err != nil {
					log.Fatalf("Failed to get items: %v", err)
				}

				// Process the response
				log.Println("List Network Interfaces:")
				for _, lif := range resp.NetInterfaces {
					log.Println("Interface with: ")
					common.PrintResponse(protojson.Format(lif))
				}

				// Are there more pages
				if resp.NextPageToken == "" {
					// if no then break
					break
				}
				// update to next token
				pageToken = resp.NextPageToken
			}
		},
	}

	cmd.Flags().Int32VarP(&pageSize, "pagesize", "s", 0, "Specify page size")
	cmd.Flags().StringVarP(&pageToken, "pagetoken", "t", "", "Specify the token")

	return cmd
}
