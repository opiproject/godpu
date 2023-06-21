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

			// create all
			obj0, err := evpnClient.CreateInterface(ctx)
			if err != nil {
				log.Fatalf("could not create interface: %v", err)
			}
			log.Printf("%s", obj0)
			obj1, err := evpnClient.CreateVpc(ctx)
			if err != nil {
				log.Fatalf("could not create vpc: %v", err)
			}
			log.Printf("%s", obj1)
			// get all
			obj2, err := evpnClient.GetInterface(ctx)
			if err != nil {
				log.Fatalf("could not get interface: %v", err)
			}
			log.Printf("%s", obj2)
			obj3, err := evpnClient.GetVpc(ctx)
			if err != nil {
				log.Fatalf("could not get vpc: %v", err)
			}
			log.Printf("%s", obj3)
			// delete all
			obj4, err := evpnClient.DeleteInterface(ctx)
			if err != nil {
				log.Fatalf("could not delete interface: %v", err)
			}
			log.Printf("%s", obj4)
			obj5, err := evpnClient.DeleteVpc(ctx)
			if err != nil {
				log.Fatalf("could not delete vpc: %v", err)
			}
			log.Printf("%s", obj5)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	return cmd
}
