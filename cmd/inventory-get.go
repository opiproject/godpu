// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package cmd implements the CLI commands
package cmd

import (
	"context"
	"log"
	"time"

	"github.com/opiproject/godpu/inventory"
	"github.com/spf13/cobra"
)

// NewGetCommand returns the inventory get command
func NewGetCommand() *cobra.Command {
	var (
		addr string
	)
	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "Gets DPU inventory information",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			invClient, err := inventory.NewClient(addr)
			if err != nil {
				log.Fatalf("could create gRPC client: %v", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			data, err := invClient.Get(ctx)
			if err != nil {
				log.Fatalf("could not get inventory: %v", err)
			}
			log.Printf("%s", data)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&addr, "addr", "localhost:50151", "address of OPI gRPC server")
	return cmd
}
