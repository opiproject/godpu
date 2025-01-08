// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2025 Dell Inc, or its subsidiaries.

// Package network implements the network related CLI commands
package netintf

import (
//	"log"
//	"time"

//	"github.com/opiproject/godpu/cmd/common"
	"github.com/spf13/cobra"
)


// ListNetInterfaces lists all Network Interface details from OPI server
func ListNetInterfaces() *cobra.Command {
	var pageSize int32
	var pageToken string

	cmd := &cobra.Command{
		Use:   "list-net-interfaces",
		Short: "List the network interfaces",
		Run: func(c *cobra.Command, _ []string) {
//			TODO: Add processing for Network Interface List
//			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

//			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
//			cobra.CheckErr(err)

//			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
//			cobra.CheckErr(err)

//			evpnClient, err := network.NewLogicalBridge(addr, tlsFiles)
//			if err != nil {
//				log.Fatalf("could not create gRPC client: %v", err)
//			}
//			defer cancel()

//			for {
//				resp, err := evpnClient.ListLogicalBridges(ctx, pageSize, pageToken)
//				if err != nil {
//					log.Fatalf("Failed to get items: %v", err)
//				}
				// Process the server response
//				log.Println("List Network Interfaces:")
//				for _, lb := range resp.NetInterfaces {
//					log.Println("Interface with: ")
//					PrintLB(lb)
//				}

				// Check if there are more pages to retrieve
//				if resp.NextPageToken == "" {
					// No more pages, break the loop
//					break
//				}
				// Update the page token for the next request
//				pageToken = resp.NextPageToken
//			}
		},
	}

	cmd.Flags().Int32VarP(&pageSize, "pagesize", "s", 0, "Specify page size")
	cmd.Flags().StringVarP(&pageToken, "pagetoken", "t", "", "Specify the token")

	return cmd
}