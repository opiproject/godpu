// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the network related CLI commands
package network

import (
	"log"

	"github.com/spf13/cobra"
)

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
	// Mgmt cli's
	cmd.AddCommand(DumpNetlinkDatabase())

	return cmd
}
