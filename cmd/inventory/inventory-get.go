// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package inventory implements inventory related CLI commands
package inventory

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
		Run: func(_ *cobra.Command, _ []string) {
			invClient, err := inventory.New(addr)
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

// NewInventoryCommand tests the  inventory
func NewInventoryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "inventory",
		Aliases: []string{"g"},
		Short:   "Tests inventory functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatalf("[ERROR] %s", err.Error())
			}
		},
	}

	cmd.AddCommand(NewGetCommand())

	return cmd
}
