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

	"github.com/opiproject/godpu/network"
	pb "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	"github.com/spf13/cobra"
)

// CreateBridgePort creates an Bridge Port an OPI server
func CreateBridgePort() *cobra.Command {
	var addr string
	var name string
	var mac string
	var bridgePortType string
	var logicalBridges []string

	cmd := &cobra.Command{
		Use:   "create-bp",
		Short: "Create a bridge port",
		Long:  "Create a BridgePort with the specified name, MAC address, type, and VLAN IDs",
		Run: func(_ *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewBridgePort(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()
			// grpc call to create the bridge port
			bridgePort, err := evpnClient.CreateBridgePort(ctx, name, mac, bridgePortType, logicalBridges)
			if err != nil {
				log.Fatalf("could not create Bridge Port: %v", err)
			}
			log.Printf("Created Bridge Port:\nname: %s \nstatus: %s\nptype: %s\nmac: %s\nbridges: %s\nComponent Status:\n%s\n", bridgePort.GetName(),
				pb.BPOperStatus_name[int32(bridgePort.GetStatus().GetOperStatus())], bridgePort.GetSpec().GetPtype(),
				bridgePort.GetSpec().GetMacAddress(), bridgePort.GetSpec().GetLogicalBridges(), PrintComponents(bridgePort.GetStatus().GetComponents()))
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().StringVar(&mac, "mac", "", "Specify the MAC address")
	cmd.Flags().StringVarP(&bridgePortType, "type", "t", "", "Specify the type (access or trunk)")
	cmd.Flags().StringSliceVar(&logicalBridges, "logicalBridges", []string{}, "Specify VLAN IDs (multiple values supported)")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

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
	var addr string
	var name string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "delete-bp",
		Short: "Delete a bridge port",
		Long:  "Delete a BridgePort with the specified name",
		Run: func(_ *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewBridgePort(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			_, err = evpnClient.DeleteBridgePort(ctx, name, allowMissing)
			if err != nil {
				log.Fatalf("DeleteBridgePort: Error occurred while deleting Bridge Port: %q", err)
			}
			log.Printf("Deleting BridgePort in process\n")
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "Specify if missing allowed")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

	return cmd
}

// GetBridgePort Get Bridge Port details
func GetBridgePort() *cobra.Command {
	var addr string
	var name string

	cmd := &cobra.Command{
		Use:   "get-bp",
		Short: "Show details of a bridge port",
		Long:  "Show details of a BridgePort with the specified name",
		Run: func(_ *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewBridgePort(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			bridgePort, err := evpnClient.GetBridgePort(ctx, name)
			if err != nil {
				log.Fatalf("GetBridgePort: Error occurred while creating Bridge Port: %q", err)
			}
			log.Printf("Get Bridge Port:\nname: %s \nstatus: %s\nptype: %s\nmac: %s\nbridges: %s\nComponent Status:\n%s\n", bridgePort.GetName(),
				pb.BPOperStatus_name[int32(bridgePort.GetStatus().GetOperStatus())], bridgePort.GetSpec().GetPtype(), bridgePort.GetSpec().GetMacAddress(),
				bridgePort.GetSpec().GetLogicalBridges(), PrintComponents(bridgePort.GetStatus().GetComponents()))
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// ListBridgePorts list all the Bridge Port an OPI server
func ListBridgePorts() *cobra.Command {
	var addr string
	var pageSize int32
	var pageToken string

	cmd := &cobra.Command{
		Use:   "list-bps",
		Short: "Show details of all bridge ports",
		Run: func(_ *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewBridgePort(addr)
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
				for _, bridgePort := range resp.BridgePorts {
					log.Printf(" Bridge Port :\nname: %s \nstatus: %s\nptype: %s\nmac: %s\nbridges: %s\nComponent Status:\n%s\n", bridgePort.GetName(),
						pb.BPOperStatus_name[int32(bridgePort.GetStatus().GetOperStatus())], bridgePort.GetSpec().GetPtype(), bridgePort.GetSpec().GetMacAddress(),
						bridgePort.GetSpec().GetLogicalBridges(), PrintComponents(bridgePort.GetStatus().GetComponents()))
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

// UpdateBridgePort update the Bridge Port on OPI server
func UpdateBridgePort() *cobra.Command {
	var addr string
	var name string
	var updateMask []string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "update-bp",
		Short: "Update the bridge port",
		Long:  "updates the Bridge Port with updated mask",
		Run: func(_ *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewBridgePort(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			bridgePort, err := evpnClient.UpdateBridgePort(ctx, name, updateMask, allowMissing)
			if err != nil {
				log.Fatalf("UpdateBridgePort: Error occurred while creating Bridge Port: %q", err)
			}
			log.Printf(" Bridge Port:\nname: %s \nstatus: %s\nptype: %s\nmac: %s\nbridges: %s\nComponent Status:\n%s\n", bridgePort.GetName(),
				pb.BPOperStatus_name[int32(bridgePort.GetStatus().GetOperStatus())], bridgePort.GetSpec().GetPtype(), bridgePort.GetSpec().GetMacAddress(),
				bridgePort.GetSpec().GetLogicalBridges(), PrintComponents(bridgePort.GetStatus().GetComponents()))
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "name of the Bridge Port")
	cmd.Flags().StringSliceVar(&updateMask, "update-mask", nil, "update mask")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "allow the missing")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	return cmd
}
