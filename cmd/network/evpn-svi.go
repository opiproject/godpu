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

// CreateSVI create svi on OPI server
func CreateSVI() *cobra.Command {
	var name string
	var vrf string
	var logicalBridge string
	var mac string
	var gwIPs []string
	var ebgp bool
	var remoteAS uint32

	cmd := &cobra.Command{
		Use:   "create-svi",
		Short: "Create a SVI",
		Long:  "Create an  using name, vrf,logical bridges, mac, gateway ip's and enable bgp ",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewSVI(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			svi, err := evpnClient.CreateSvi(ctx, name, vrf, logicalBridge, mac, gwIPs, ebgp, remoteAS)
			if err != nil {
				log.Fatalf("failed to create logical bridge: %v", err)
			}

			log.Println("Created SVI:")
			PrintSvi(svi)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "SVI Name")
	cmd.Flags().StringVar(&vrf, "vrf", "", "Must be unique")
	cmd.Flags().StringVar(&logicalBridge, "logical-bridge", "", "Pair of vni and vlan_id must be unique")
	cmd.Flags().StringVar(&mac, "mac", "", "GW MAC address, random MAC assigned if not specified")
	cmd.Flags().StringSliceVar(&gwIPs, "gw-ips", nil, "List of GW IP addresses")
	cmd.Flags().BoolVar(&ebgp, "ebgp", false, "Enable eBGP in VRF for tenants connected through this SVI")
	cmd.Flags().Uint32VarP(&remoteAS, "remote-as", "", 0, "The remote AS")

	if err := cmd.MarkFlagRequired("vrf"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	if err := cmd.MarkFlagRequired("logical-bridge"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	if err := cmd.MarkFlagRequired("mac"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	if err := cmd.MarkFlagRequired("gw-ips"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// DeleteSVI delete the svi on OPI server
func DeleteSVI() *cobra.Command {
	var name string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "delete-svi",
		Short: "Delete a SVI",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewSVI(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			_, err = evpnClient.DeleteSvi(ctx, name, allowMissing)
			if err != nil {
				log.Fatalf("failed to create logical bridge: %v", err)
			}

			log.Printf("Deleted SVI: %s\n", name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().BoolVarP(&allowMissing, "allow-missing", "a", false, "Specify the name of the BridgePort")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// GetSVI get svi details from OPI server
func GetSVI() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "get-svi",
		Short: "Show details of a SVI",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewSVI(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			svi, err := evpnClient.GetSvi(ctx, name)
			if err != nil {
				log.Fatalf("GetSVI: Error occurred while creating Bridge Port: %q", err)
			}
			log.Println("Get SVI:")
			PrintSvi(svi)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// ListSVIs get all the svi's from OPI server
func ListSVIs() *cobra.Command {
	var pageSize int32
	var pageToken string

	cmd := &cobra.Command{
		Use:   "list-svis",
		Short: "Show details of all SVIs",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewSVI(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()
			for {
				resp, err := evpnClient.ListSvis(ctx, pageSize, pageToken)
				if err != nil {
					log.Fatalf("Failed to get items: %v", err)
				}
				// Process the server response
				log.Println("List SVIs:")
				for _, svi := range resp.Svis {
					log.Println("SVI with:")
					PrintSvi(svi)
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

	cmd.Flags().Int32VarP(&pageSize, "pageSize", "s", 0, "Specify the name of the BridgePort")
	cmd.Flags().StringVarP(&pageToken, "pageToken", "p", "", "Specify the page token")

	return cmd
}

// UpdateSVI update the svi on OPI server
func UpdateSVI() *cobra.Command {
	var name string
	var updateMask []string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "update-svi",
		Short: "update the SVI",
		Run: func(c *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			evpnClient, err := network.NewSVI(addr, tlsFiles)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			svi, err := evpnClient.UpdateSvi(ctx, name, updateMask, allowMissing)
			if err != nil {
				log.Fatalf("GetBridgePort: Error occurred while creating Bridge Port: %q", err)
			}
			log.Println("Updated SVI:")
			PrintSvi(svi)
		},
	}
	cmd.Flags().StringSliceVar(&updateMask, "update-mask", nil, "update mask")
	cmd.Flags().BoolVarP(&allowMissing, "allow-missing", "a", false, "allow the missing")

	return cmd
}
