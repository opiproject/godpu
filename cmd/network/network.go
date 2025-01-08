// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2024 Dell Inc, or its subsidiaries.

// Package network implements the network related CLI commands
package network

import (
	"log"
	"time"

	"github.com/opiproject/godpu/cmd/common"
	"github.com/opiproject/godpu/cmd/network/netintf"
	"github.com/opiproject/godpu/cmd/network/evpn"
	"github.com/spf13/cobra"
)

// NewNetworkCommand tests the Network functionality command
func NewNetworkCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "network",
		Aliases: []string{"g"},
		Short:   "Tests DPU networking functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatalf("[ERROR] %s", err.Error())
			}
		},
	}

	flags := cmd.PersistentFlags()
	flags.Duration(common.TimeoutCmdLineArg, 10*time.Second, "timeout for a cmd")

	cmd.AddCommand(NewEvpnCommand())
	cmd.AddCommand(NewNetIntfCommand())

	return cmd
}

// NewNetworkCommand tests the Network functionality command
func NewNetIntfCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "intf",
		Aliases: []string{"g"},
		Short:   "Tests DPU network interface functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatalf("[ERROR] %s", err.Error())
			}
		},
	}

	flags := cmd.PersistentFlags()
	flags.Duration(common.TimeoutCmdLineArg, 10*time.Second, "timeout for a cmd")

	cmd.AddCommand(netintf.ListNetInterfaces())

	return cmd
}

// NewEvpnCommand tests the EVPN functionality command
func NewEvpnCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "evpn",
		Aliases: []string{"g"},
		Short:   "Tests DPU evpn functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatalf("[ERROR] %s", err.Error())
			}
		},
	}
	// Bridge cli's
	cmd.AddCommand(evpn.CreateLogicalBridge())
	cmd.AddCommand(evpn.DeleteLogicalBridge())
	cmd.AddCommand(evpn.GetLogicalBridge())
	cmd.AddCommand(evpn.ListLogicalBridges())
	cmd.AddCommand(evpn.UpdateLogicalBridge())
	// Port cli's
	cmd.AddCommand(evpn.CreateBridgePort())
	cmd.AddCommand(evpn.DeleteBridgePort())
	cmd.AddCommand(evpn.GetBridgePort())
	cmd.AddCommand(evpn.ListBridgePorts())
	cmd.AddCommand(evpn.UpdateBridgePort())
	// VRF cli's
	cmd.AddCommand(evpn.CreateVRF())
	cmd.AddCommand(evpn.DeleteVRF())
	cmd.AddCommand(evpn.GetVRF())
	cmd.AddCommand(evpn.ListVRFs())
	cmd.AddCommand(evpn.UpdateVRF())
	// SVI cli's
	cmd.AddCommand(evpn.CreateSVI())
	cmd.AddCommand(evpn.DeleteSVI())
	cmd.AddCommand(evpn.GetSVI())
	cmd.AddCommand(evpn.ListSVIs())
	cmd.AddCommand(evpn.UpdateSVI())

	return cmd
}
