// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package cmd implements the CLI commands
package cmd

import (
	"context"
	"log"
	"time"

	"github.com/opiproject/godpu/network"
	"github.com/spf13/cobra"
)

// CreateLogicalBridge creates an Logical Bridge an OPI server
func CreateLogicalBridge() *cobra.Command {
	var addr string
	var name string
	var vlanID uint32
	var vni uint32

	cmd := &cobra.Command{
		Use:   "create-lb",
		Short: "Create a logical bridge",
		Long:  "Create a logical bridge with the specified name, VLAN ID, and VNI",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewLogicalBridge(addr)
			if err != nil {
				log.Fatalf("could create gRPC client: %v", err)
			}
			defer cancel()

			resp, err := evpnClient.CreateLogicalBridge(ctx, name, vlanID, vni)
			if err != nil {
				log.Fatalf("failed to create logical bridge: %v", err)
			}

			log.Printf(" Created Logical Bridge \n name: %s,\n vlan: %d,\n vni: %d,\n status: %s\n", resp.GetName(), resp.GetSpec().GetVlanId(),
				resp.GetSpec().GetVni(), resp.GetStatus())
		},
	}
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the logical bridge")
	cmd.Flags().Uint32VarP(&vlanID, "vlan-id", "v", 0, "Specify the VLAN ID")
	cmd.Flags().Uint32VarP(&vni, "vni", "i", 0, "Specify the VNI")

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
		Run: func(cmd *cobra.Command, args []string) {
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

			log.Printf("Deleted Logical Bridge: %s\n", resp)
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
		Run: func(cmd *cobra.Command, args []string) {
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

			log.Printf(" Created Logical Bridge \n name: %s,\n vlan: %d,\n vni: %d,\n status: %s\n", resp.GetName(), resp.GetSpec().GetVlanId(),
				resp.GetSpec().GetVni(), resp.GetStatus())
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
		Run: func(cmd *cobra.Command, args []string) {
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
					log.Printf(" Created Logical Bridge \n name: %s,\n vlan: %d,\n vni: %d,\n status: %s\n", item.GetName(), item.GetSpec().GetVlanId(),
						item.GetSpec().GetVni(), item.GetStatus())
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
		Run: func(cmd *cobra.Command, args []string) {
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

			log.Printf(" Updated Logical Bridge \n name: %s,\n vlan: %d,\n vni: %d,\n status: %s\n", resp.GetName(), resp.GetSpec().GetVlanId(),
				resp.GetSpec().GetVni(), resp.GetStatus())
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "name of the logical bridge")
	cmd.Flags().StringSliceVar(&updateMask, "update-mask", nil, "update mask")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

	return cmd
}

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
		Run: func(cmd *cobra.Command, args []string) {
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
			log.Printf("Created Bridge Port:\n status: %s\n, type: %s\n, name: %s\n, bridges: %s\n, mac: %s\n", bridgePort.GetStatus().GetOperStatus(), bridgePort.GetSpec().GetPtype(),
				bridgePort.GetName(), bridgePort.GetSpec().GetLogicalBridges(), bridgePort.GetSpec().GetMacAddress())
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
	err := cmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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
		Run: func(cmd *cobra.Command, args []string) {
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
			log.Printf("Deleted BridgePort ")
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
		Run: func(cmd *cobra.Command, args []string) {
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
			log.Printf("Bridge Port:\n status: %s\n, type: %s\n, name: %s\n, bridges: %s\n, mac: %s\n", bridgePort.GetStatus().GetOperStatus(), bridgePort.GetSpec().GetPtype(),
				bridgePort.GetName(), bridgePort.GetSpec().GetLogicalBridges(), bridgePort.GetSpec().GetMacAddress())
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
		Run: func(cmd *cobra.Command, args []string) {
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
					log.Printf("Bridge Port:\n status: %s\n, type: %s\n, name: %s\n, bridges: %s\n, mac: %s\n", bridgePort.GetStatus().GetOperStatus(), bridgePort.GetSpec().GetPtype(),
						bridgePort.GetName(), bridgePort.GetSpec().GetLogicalBridges(), bridgePort.GetSpec().GetMacAddress())
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
		Run: func(cmd *cobra.Command, args []string) {
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
			log.Printf("Bridge Port:\n status: %s\n, type: %s\n, name: %s\n, bridges: %s\n, mac: %s\n", bridgePort.GetStatus().GetOperStatus(), bridgePort.GetSpec().GetPtype(),
				bridgePort.GetName(), bridgePort.GetSpec().GetLogicalBridges(), bridgePort.GetSpec().GetMacAddress())
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "name of the Bridge Port")
	cmd.Flags().StringSliceVar(&updateMask, "update-mask", nil, "update mask")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "allow the missing")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	return cmd
}

// CreateVRF Create vrf on OPI Server
func CreateVRF() *cobra.Command {
	var addr string
	var name string
	var vni uint32
	var loopback string
	var vtep string

	cmd := &cobra.Command{
		Use:   "create-vrf",
		Short: "Create a VRF",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewVRF(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			vrf, err := evpnClient.CreateVrf(ctx, name, vni, loopback, vtep)
			if err != nil {
				log.Fatalf("failed to create logical bridge: %v", err)
			}
			log.Printf("Created VRF with name: %s\n, vni : %d\n, vtep ip : %s\n, loopback ip: %s\n", vrf.GetName(), vrf.GetSpec().GetVni(),
				vrf.GetSpec().GetVtepIpPrefix(), vrf.GetSpec().GetLoopbackIpPrefix())
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Descriptive name")
	cmd.Flags().Uint32VarP(&vni, "vni", "v", 0, "Must be unique ")
	cmd.Flags().StringVar(&loopback, "loopback", "", "Loopback IP address")
	cmd.Flags().StringVar(&vtep, "vtep", "", "VTEP IP address")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

	if err := cmd.MarkFlagRequired("vni"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	if err := cmd.MarkFlagRequired("loopback"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// DeleteVRF update the vrf on OPI server
func DeleteVRF() *cobra.Command {
	var addr string
	var name string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "delete-vrf",
		Short: "Delete a VRF",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewVRF(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			_, err = evpnClient.DeleteVrf(ctx, name, allowMissing)
			if err != nil {
				log.Fatalf("DeleteVRF: Error occurred while creating Bridge Port: %q", err)
			}
			log.Printf("Deleted VRF with VPort ID: %s\n", name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "Specify the name of the BridgePort")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

	return cmd
}

// GetVRF get vrf details from OPI server
func GetVRF() *cobra.Command {
	var addr string
	var name string

	cmd := &cobra.Command{
		Use:   "get-vrf",
		Short: "Show details of a VRF",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewVRF(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			vrf, err := evpnClient.GetVrf(ctx, name)
			if err != nil {
				log.Fatalf("DeleteVRF: Error occurred while creating Bridge Port: %q", err)
			}
			log.Printf("VRF with name: %s\n, operation status: %d\n,vni : %d\n, vtep ip : %s\n, loopback ip: %s\n", vrf.GetName(), vrf.GetStatus().GetOperStatus(),
				vrf.GetSpec().GetVni(), vrf.GetSpec().GetVtepIpPrefix(), vrf.GetSpec().GetLoopbackIpPrefix())
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the vrf")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// ListVRFs list all vrf's with details from OPI server
func ListVRFs() *cobra.Command {
	var addr string
	var pageSize int32
	var pageToken string

	cmd := &cobra.Command{
		Use:   "list-vrfs",
		Short: "Show details of all Vrfs",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewVRF(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()
			for {
				resp, err := evpnClient.ListVrfs(ctx, pageSize, pageToken)
				if err != nil {
					log.Fatalf("Failed to get items: %v", err)
				}
				// Process the server response
				for _, vrf := range resp.Vrfs {
					log.Printf("VRF with name: %s\n, operation status: %d\n,vni : %d\n, vtep ip : %s\n, loopback ip: %s\n", vrf.GetName(), vrf.GetStatus().GetOperStatus(),
						vrf.GetSpec().GetVni(), vrf.GetSpec().GetVtepIpPrefix(), vrf.GetSpec().GetLoopbackIpPrefix())
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

// UpdateVRF update the vrf on OPI server
func UpdateVRF() *cobra.Command {
	var addr string
	var name string
	var updateMask []string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "update-vrf",
		Short: "update the VRF",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewVRF(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			vrf, err := evpnClient.UpdateVrf(ctx, name, updateMask, allowMissing)
			if err != nil {
				log.Fatalf("GetBridgePort: Error occurred while creating Bridge Port: %q", err)
			}
			log.Printf("Updated VRF with name: %s\n, operation status: %d\n,vni : %d\n, vtep ip : %s\n, loopback ip: %s\n", vrf.GetName(), vrf.GetStatus().GetOperStatus(),
				vrf.GetSpec().GetVni(), vrf.GetSpec().GetVtepIpPrefix(), vrf.GetSpec().GetLoopbackIpPrefix())
		},
	}
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the vrf")
	cmd.Flags().StringSliceVar(&updateMask, "update-mask", nil, "update mask")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "allow the missing")

	return cmd
}

// CreateSVI create svi on OPI server
func CreateSVI() *cobra.Command {
	var addr string
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
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewSVI(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			svi, err := evpnClient.CreateSvi(ctx, name, vrf, logicalBridge, mac, gwIPs, ebgp, remoteAS)
			if err != nil {
				log.Fatalf("failed to create logical bridge: %v", err)
			}

			log.Printf("CreateSVI: Created SVI  \n name: %s,\n status: %d,\n Vrf: %s,\n LogicalBridge: %s,\n MacAddress: %s,\n EnableBgp: %t,\n GwIPs: %s,\nremoteAS: %d\n",
				svi.GetName(), svi.GetStatus().GetOperStatus(), svi.GetSpec().GetVrf(), svi.GetSpec().GetLogicalBridge(), svi.GetSpec().GetMacAddress(),
				svi.GetSpec().GetEnableBgp(), svi.GetSpec().GetGwIpPrefix(), svi.GetSpec().GetRemoteAs())
		},
	}
	cmd.Flags().StringVar(&vrf, "vrf", "", "Must be unique")
	cmd.Flags().StringVar(&logicalBridge, "logicalBridge", "", "Pair of vni and vlan_id must be unique")
	cmd.Flags().StringVar(&mac, "mac", "", "GW MAC address, random MAC assigned if not specified")
	cmd.Flags().StringSliceVar(&gwIPs, "gw-ips", nil, "List of GW IP addresses")
	cmd.Flags().BoolVar(&ebgp, "ebgp", false, "Enable eBGP in VRF for tenants connected through this SVI")
	cmd.Flags().Uint32VarP(&remoteAS, "remote-as", "", 0, "The remote AS")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

	if err := cmd.MarkFlagRequired("vrf"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	if err := cmd.MarkFlagRequired("logicalBridge"); err != nil {
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
	var addr string
	var name string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "delete-svi",
		Short: "Delete a SVI",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewSVI(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			_, err = evpnClient.DeleteSvi(ctx, name, allowMissing)
			if err != nil {
				log.Fatalf("failed to create logical bridge: %v", err)
			}

			log.Printf("Deleted SVI ")
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "Specify the name of the BridgePort")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// GetSVI get svi details from OPI server
func GetSVI() *cobra.Command {
	var addr string
	var name string

	cmd := &cobra.Command{
		Use:   "get-svi",
		Short: "Show details of a SVI",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewSVI(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			svi, err := evpnClient.GetSvi(ctx, name)
			if err != nil {
				log.Fatalf("GetSVI: Error occurred while creating Bridge Port: %q", err)
			}
			log.Printf("GetSVI: Created SVI  \n name: %s,\n status: %d,\n Vrf: %s,\n LogicalBridge: %s,\n MacAddress: %s,\n EnableBgp: %t,\n GwIPs: %s,\nremoteAS: %d\n",
				svi.GetName(), svi.GetStatus().GetOperStatus(), svi.GetSpec().GetVrf(), svi.GetSpec().GetLogicalBridge(), svi.GetSpec().GetMacAddress(),
				svi.GetSpec().GetEnableBgp(), svi.GetSpec().GetGwIpPrefix(), svi.GetSpec().GetRemoteAs())
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the BridgePort")
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// ListSVIs get all the svi's from OPI server
func ListSVIs() *cobra.Command {
	var addr string
	var pageSize int32
	var pageToken string

	cmd := &cobra.Command{
		Use:   "list-svis",
		Short: "Show details of all SVIs",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewSVI(addr)
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
				for _, svi := range resp.Svis {
					log.Printf("ListSVIs: SVI  \n name: %s,\n status: %d,\n Vrf: %s,\n LogicalBridge: %s,\n MacAddress: %s,\n EnableBgp: %t,\n GwIPs: %s,\nremoteAS: %d\n",
						svi.GetName(), svi.GetStatus().GetOperStatus(), svi.GetSpec().GetVrf(), svi.GetSpec().GetLogicalBridge(), svi.GetSpec().GetMacAddress(),
						svi.GetSpec().GetEnableBgp(), svi.GetSpec().GetGwIpPrefix(), svi.GetSpec().GetRemoteAs())
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
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	return cmd
}

// UpdateSVI update the svi on OPI server
func UpdateSVI() *cobra.Command {
	var addr string
	var name string
	var updateMask []string
	var allowMissing bool

	cmd := &cobra.Command{
		Use:   "update-svi",
		Short: "update the SVI",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			evpnClient, err := network.NewSVI(addr)
			if err != nil {
				log.Fatalf("could not create gRPC client: %v", err)
			}
			defer cancel()

			// grpc call to create the bridge port
			svi, err := evpnClient.UpdateSvi(ctx, name, updateMask, allowMissing)
			if err != nil {
				log.Fatalf("GetBridgePort: Error occurred while creating Bridge Port: %q", err)
			}
			log.Printf("UpdateSVI: SVI  \n name: %s,\n status: %d,\n Vrf: %s,\n LogicalBridge: %s,\n MacAddress: %s,\n EnableBgp: %t,\n GwIPs: %s,\nremoteAS: %d\n",
				svi.GetName(), svi.GetStatus().GetOperStatus(), svi.GetSpec().GetVrf(), svi.GetSpec().GetLogicalBridge(), svi.GetSpec().GetMacAddress(),
				svi.GetSpec().GetEnableBgp(), svi.GetSpec().GetGwIpPrefix(), svi.GetSpec().GetRemoteAs())
		},
	}
	cmd.Flags().StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	cmd.Flags().StringSliceVar(&updateMask, "update-mask", nil, "update mask")
	cmd.Flags().BoolVarP(&allowMissing, "allowMissing", "a", false, "allow the missing")
	return cmd
}

// NewEvpnCommand tests the EVPN functionality command
func NewEvpnCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "evpn",
		Aliases: []string{"g"},
		Short:   "Tests DPU evpn functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatalf("[ERROR] %s", err.Error())
			}
		},
	}
	// Bridge cli's
	cmd.AddCommand(CreateLogicalBridge())
	cmd.AddCommand(DeleteLogicalBridge())
	cmd.AddCommand(GetLogicalBridge())
	cmd.AddCommand(ListLogicalBridges())
	cmd.AddCommand(UpdateLogicalBridge())
	// Port cli's
	cmd.AddCommand(CreateBridgePort())
	cmd.AddCommand(DeleteBridgePort())
	cmd.AddCommand(GetBridgePort())
	cmd.AddCommand(ListBridgePorts())
	cmd.AddCommand(UpdateBridgePort())
	// VRF cli's
	cmd.AddCommand(CreateVRF())
	cmd.AddCommand(DeleteVRF())
	cmd.AddCommand(GetVRF())
	cmd.AddCommand(ListVRFs())
	cmd.AddCommand(UpdateVRF())
	// SVI cli's
	cmd.AddCommand(CreateSVI())
	cmd.AddCommand(DeleteSVI())
	cmd.AddCommand(GetSVI())
	cmd.AddCommand(ListSVIs())
	cmd.AddCommand(UpdateSVI())

	return cmd
}
