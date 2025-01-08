// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (c) 2024 Ericsson AB.

// Package evpn implements the evpn network related CLI commands
package evpn

import (
	"context"
	"log"
	"time"

	"github.com/opiproject/godpu/cmd/common"
	"github.com/opiproject/godpu/network"
	"github.com/spf13/cobra"
)

// CreateLogicalBridge creates an Logical Bridge an OPI server
func CreateLogicalBridge() *cobra.Command {
	var name string
	var vlanID uint32
	var vni uint32
	var vtep string

	cmd := &cobra.Command{
		Use:   "create-lb",
		Short: "Create a logical bridge",
		Long:  "Create a logical bridge with the specified name, VLAN ID, and VNI",
		Run: func(c *cobra.Command, _ []string) {
			var vniparam *uint32
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewLogicalBridge(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could create gRPC client: %v", err)
			}
			defer cancel()

			if vni != 0 {
				vniparam = &vni
			}

			lb, err := evpnClient.CreateLogicalBridge(ctx, name, vlanID, vniparam, vtep)
			if err != nil {
				log.Fatalf("failed to create logical bridge: %v", err)
			}

			log.Println("Created Logical Bridge:")
			PrintLB(lb)
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the logical bridge")
	cmd.Flags().Uint32VarP(&vlanID, "vlan-id", "v", 0, "Specify the VLAN ID")
	cmd.Flags().Uint32VarP(&vni, "vni", "i", 0, "Specify the VNI")
	cmd.Flags().StringVar(&vtep, "vtep", "", "VTEP IP address")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("vlan-id"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	cmd.MarkFlagsRequiredTogether("vni", "vtep")

	return cmd
}

// DeleteLogicalBridge deletes an Logical Bridge an OPI server
func DeleteLogicalBridge() *cobra.Command {
	var name string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "delete-lb",
		Short: "Delete a logical bridge",
		Long:  "Delete a logical bridge with the specified name",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewLogicalBridge(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			_, err = evpnClient.DeleteLogicalBridge(ctx, name, allowMissing)
			if err != nil {
				log.Fatalf("failed to delete logical bridge: %v", err)
			}

			log.Printf("Deleted Logical Bridge: %s\n", name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "Specify allow missing")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// GetLogicalBridge get Logical Bridge details from OPI server
func GetLogicalBridge() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "get-lb",
		Short: "Show details of a logical bridge",
		Long:  "Show details of a logical bridge with the specified name",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewLogicalBridge(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			lb, err := evpnClient.GetLogicalBridge(ctx, name)
			if err != nil {
				log.Fatalf("failed to get logical bridge: %v", err)
			}

			log.Println("Get Logical Bridge:")
			PrintLB(lb)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	return cmd
}

// ListLogicalBridges list all Logical Bridge details from OPI server
func ListLogicalBridges() *cobra.Command {
	var pageSize int32
	var pageToken string

	cmd := &cobra.Command{
		Use:   "list-lbs",
		Short: "Show details of all logical bridges",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewLogicalBridge(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			for {
				resp, err := evpnClient.ListLogicalBridges(ctx, pageSize, pageToken)
				if err != nil {
					log.Fatalf("Failed to get items: %v", err)
				}
				// Process the server response
				log.Println("List Logical Bridges:")
				for _, lb := range resp.LogicalBridges {
					log.Println("Logical Bridge with: ")
					PrintLB(lb)
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

// UpdateLogicalBridge update Logical Bridge on OPI server
func UpdateLogicalBridge() *cobra.Command {
	var name string
	var allowMissing bool
	var updateMask []string

	cmd := &cobra.Command{
		Use:   "update-lb",
		Short: "update the logical bridge",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewLogicalBridge(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			lb, err := evpnClient.UpdateLogicalBridge(ctx, name, updateMask, allowMissing)
			if err != nil {
				log.Fatalf("failed to update logical bridge: %v", err)
			}
			log.Println("Updated Logical Bridge:")
			PrintLB(lb)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "name of the logical bridge")
	cmd.Flags().StringSliceVar(&updateMask, "update-mask", nil, "update mask")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "Specify allow missing")

	return cmd
}
