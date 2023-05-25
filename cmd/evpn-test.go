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

// NewEvpnCommand tests the EVPN functionality command
func NewEvpnCommand() *cobra.Command {
	var (
		addr string
	)
	cmd := &cobra.Command{
		Use:     "evpn",
		Aliases: []string{"g"},
		Short:   "Tests DPU evpn functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			evpnClient, err := network.New(addr)
			if err != nil {
				log.Fatalf("could create gRPC client: %v", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			data, err := evpnClient.CreateInterface(ctx)
			if err != nil {
				log.Fatalf("could not test evpn: %v", err)
			}
			log.Printf("%s", data)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	return cmd
}
