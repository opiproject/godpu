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

// CreateBridgePort creates an Bridge Port an OPI server
func CreateBridgePort() *cobra.Command {
	var name string
	var mac string
	var bridgePortType string
	var logicalBridges []string

	cmd := &cobra.Command{
		Use:   "create-bp",
		Short: "Create a bridge port",
		Long:  "Create a BridgePort with the specified name, MAC address, type, and VLAN IDs",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)
			evpnClient, err := network.NewBridgePort(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()
			// grpc call to create the bridge port
			bridgePort, err := evpnClient.CreateBridgePort(ctx, name, mac, bridgePortType, logicalBridges)
			if err != nil {
				log.Fatalf("could not create Bridge Port: %v", err)
			}

			log.Println("Created Bridge Port:")
			PrintBP(bridgePort)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().StringVar(&mac, "mac", "", "Specify the MAC address")
	cmd.Flags().StringVarP(&bridgePortType, "type", "t", "", "Specify the type (access or trunk)")
	cmd.Flags().StringSliceVar(&logicalBridges, "logicalBridges", []string{}, "Specify VLAN IDs (multiple values supported)")

	if err := cmd.MarkFlagRequired("mac"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("type"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	// Define allowed choices for the "type" Flag
	err := cmd.RegisterFlagCompletionFunc("type", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"access", "trunk"}, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		log.Fatalf("Error registering flag completion function: %v", err)
	}
	return cmd
}

// DeleteBridgePort delete an Bridge Port an OPI server
func DeleteBridgePort() *cobra.Command {
	var name string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "delete-bp",
		Short: "Delete a bridge port",
		Long:  "Delete a BridgePort with the specified name",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewBridgePort(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			_, err = evpnClient.DeleteBridgePort(ctx, name, allowMissing)
			if err != nil {
				log.Fatalf("DeleteBridgePort: Error occurred while deleting Bridge Port: %q", err)
			}
			log.Printf("Deleted BridgePort: %s\n", name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "Specify if missing allowed")

	return cmd
}

// GetBridgePort Get Bridge Port details
func GetBridgePort() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "get-bp",
		Short: "Show details of a bridge port",
		Long:  "Show details of a BridgePort with the specified name",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewBridgePort(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			bridgePort, err := evpnClient.GetBridgePort(ctx, name)
			if err != nil {
				log.Fatalf("GetBridgePort: Error occurred while creating Bridge Port: %q", err)
			}
			log.Println("Get Bridge Port:")
			PrintBP(bridgePort)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// ListBridgePorts list all the Bridge Port an OPI server
func ListBridgePorts() *cobra.Command {
	var pageSize int32
	var pageToken string

	cmd := &cobra.Command{
		Use:   "list-bps",
		Short: "Show details of all bridge ports",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewBridgePort(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()
			for {
				resp, err := evpnClient.ListBridgePorts(ctx, pageSize, pageToken)
				if err != nil {
					log.Fatalf("Failed to get items: %v", err)
				}
				// Process the server response
				log.Println("List Bridge Ports:")
				for _, bridgePort := range resp.BridgePorts {
					log.Println("Bridge Port with:")
					PrintBP(bridgePort)
				}

				// Check if there are more pages to retrieve
				if resp.NextPageToken == "" {
					// No more pages, break the loop
					break
				}
				// Update the page token for the next request
				pageToken = resp.NextPageToken
			}
		},
	}
	cmd.Flags().Int32VarP(&pageSize, "pagesize", "s", 0, "Specify page size")
	cmd.Flags().StringVarP(&pageToken, "pagetoken", "t", "", "Specify the token")

	return cmd
}

// UpdateBridgePort update the Bridge Port on OPI server
func UpdateBridgePort() *cobra.Command {
	var name string
	var updateMask []string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "update-bp",
		Short: "Update the bridge port",
		Long:  "updates the Bridge Port with updated mask",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewBridgePort(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			bridgePort, err := evpnClient.UpdateBridgePort(ctx, name, updateMask, allowMissing)
			if err != nil {
				log.Fatalf("UpdateBridgePort: Error occurred while creating Bridge Port: %q", err)
			}

			log.Println("Updated Bridge Port:")
			PrintBP(bridgePort)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "name of the Bridge Port")
	cmd.Flags().StringSliceVar(&updateMask, "update-mask", nil, "update mask")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "allow the missing")

	return cmd
}
