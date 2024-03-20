// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (c) 2024 Ericsson AB.

// Package network implements the network related CLI commands
package network

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PraserX/ipconv"
	"github.com/opiproject/godpu/network"
	pb "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	"github.com/spf13/cobra"
)

// CreateLogicalBridge creates an Logical Bridge an OPI server
func CreateLogicalBridge() *cobra.Command {
	var addr string
	var name string
	var vlanID uint32
	var vni uint32
	var vtep string

	cmd := &cobra.Command{
		Use:   "create-lb",
		Short: "Create a logical bridge",
		Long:  "Create a logical bridge with the specified name, VLAN ID, and VNI",
		Run: func(_ *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewLogicalBridge(addr)
			if err != nil {
				log.Fatalf("could create gRPC client: %v", err)
			}
			defer cancel()

			resp, err := evpnClient.CreateLogicalBridge(ctx, name, vlanID, vni, vtep)
			if err != nil {
				log.Fatalf("failed to create logical bridge: %v", err)
			}
			Vteip := fmt.Sprintf("%+v/%v", ipconv.IntToIPv4(resp.GetSpec().GetVtepIpPrefix().GetAddr().GetV4Addr()), resp.GetSpec().GetVtepIpPrefix().GetLen())
			log.Printf(" Created Logical Bridge \nname: %s\nstatus: %s\nvlan: %d\nvni: %d\nVtepIpPrefix:%s\nComponent Status:\n%s\n", resp.GetName(),
				pb.LBOperStatus_name[int32(resp.GetStatus().GetOperStatus())], resp.GetSpec().GetVlanId(), resp.GetSpec().GetVni(), Vteip, PrintComponents(resp.GetStatus().GetComponents()))
		},
	}
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the logical bridge")
	cmd.Flags().Uint32VarP(&vlanID, "vlan-id", "v", 0, "Specify the VLAN ID")
	cmd.Flags().Uint32VarP(&vni, "vni", "i", 0, "Specify the VNI")
	cmd.Flags().StringVar(&vtep, "vtep", "", "VTEP IP address")

	if err := cmd.MarkFlagRequired("addr"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("vlan-id"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("vni"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	return cmd
}

// DeleteLogicalBridge deletes an Logical Bridge an OPI server
func DeleteLogicalBridge() *cobra.Command {
	var addr string
	var name string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "delete-lb",
		Short: "Delete a logical bridge",
		Long:  "Delete a logical bridge with the specified name",
		Run: func(_ *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewLogicalBridge(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			resp, err := evpnClient.DeleteLogicalBridge(ctx, name, allowMissing)
			if err != nil {
				log.Fatalf("failed to delete logical bridge: %v", err)
			}

			log.Printf("Deleting Logical Bridge in process: %s\n", resp)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "Specify allow missing")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// GetLogicalBridge get Logical Bridge details from OPI server
func GetLogicalBridge() *cobra.Command {
	var addr string
	var name string

	cmd := &cobra.Command{
		Use:   "get-lb",
		Short: "Show details of a logical bridge",
		Long:  "Show details of a logical bridge with the specified name",
		Run: func(_ *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewLogicalBridge(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			resp, err := evpnClient.GetLogicalBridge(ctx, name)
			if err != nil {
				log.Fatalf("failed to get logical bridge: %v", err)
			}
			Vteip := fmt.Sprintf("%+v/%v", ipconv.IntToIPv4(resp.GetSpec().GetVtepIpPrefix().GetAddr().GetV4Addr()), resp.GetSpec().GetVtepIpPrefix().GetLen())
			log.Printf(" Get Logical Bridge \nname: %s\nstatus: %s\nvlan: %d\nvni: %d\nVtepIpPrefix:%s\nComponent Status:\n%s\n", resp.GetName(),
				pb.LBOperStatus_name[int32(resp.GetStatus().GetOperStatus())], resp.GetSpec().GetVlanId(), resp.GetSpec().GetVni(), Vteip, PrintComponents(resp.GetStatus().GetComponents()))
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	return cmd
}

// ListLogicalBridges list all Logical Bridge details from OPI server
func ListLogicalBridges() *cobra.Command {
	var addr string
	var pageSize int32
	var pageToken string
	cmd := &cobra.Command{
		Use:   "list-lbs",
		Short: "Show details of all logical bridges",
		Run: func(_ *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewLogicalBridge(addr)
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
				for _, item := range resp.LogicalBridges {
					Vteip := fmt.Sprintf("%+v/%v", ipconv.IntToIPv4(item.GetSpec().GetVtepIpPrefix().GetAddr().GetV4Addr()), item.GetSpec().GetVtepIpPrefix().GetLen())
					log.Printf("Logical Bridge with \nname: %s\nstatus: %s\nvlan: %d\nvni: %d\nVtepIpPrefix:%s\nComponent Status:\n%s\n", item.GetName(),
						pb.LBOperStatus_name[int32(item.GetStatus().GetOperStatus())], item.GetSpec().GetVlanId(), item.GetSpec().GetVni(), Vteip, PrintComponents(item.GetStatus().GetComponents()))
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
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	return cmd
}

// UpdateLogicalBridge update Logical Bridge on OPI server
func UpdateLogicalBridge() *cobra.Command {
	var addr string
	var name string
	var updateMask []string
	cmd := &cobra.Command{
		Use:   "update-lb",
		Short: "update the logical bridge",
		Run: func(_ *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewLogicalBridge(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			resp, err := evpnClient.UpdateLogicalBridge(ctx, name, updateMask)
			if err != nil {
				log.Fatalf("failed to update logical bridge: %v", err)
			}
			Vteip := fmt.Sprintf("%+v/%v", ipconv.IntToIPv4(resp.GetSpec().GetVtepIpPrefix().GetAddr().GetV4Addr()), resp.GetSpec().GetVtepIpPrefix().GetLen())
			log.Printf(" Updated Logical Bridge with \nname: %s\nstatus: %s\nvlan: %d\nvni: %d\nVtepIpPrefix:%s\nComponent Status:\n%s\n", resp.GetName(),
				pb.LBOperStatus_name[int32(resp.GetStatus().GetOperStatus())], resp.GetSpec().GetVlanId(), resp.GetSpec().GetVni(), Vteip, PrintComponents(resp.GetStatus().GetComponents()))
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "name of the logical bridge")
	cmd.Flags().StringSliceVar(&updateMask, "update-mask", nil, "update mask")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

	return cmd
}
