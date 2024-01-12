// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package cmd implements the CLI commands
package cmd

import (
	"context"

	"github.com/opiproject/godpu/storage"
	"github.com/spf13/cobra"
)

func newCreateVirtioBlkCommand() *cobra.Command {
	id := ""
	volume := ""
	var port uint
	var pf uint
	var vf uint
	var maxIoQPS uint
	cmd := &cobra.Command{
		Use:     "blk",
		Aliases: []string{"b"},
		Short:   "Creates virtio-blk controller",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			addr, err := c.Flags().GetString(addrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(timeoutCmdLineArg)
			cobra.CheckErr(err)

			client, err := storage.New(addr)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			response, err := client.CreateVirtioBlk(ctx, id, volume, port, pf, vf, maxIoQPS)
			cobra.CheckErr(err)

			printResponse(response.Name)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "id for created resource. Assigned by server if omitted.")
	cmd.Flags().StringVar(&volume, "volume", "", "volume name to attach to virtio-blk controller")
	cmd.Flags().UintVar(&port, "port", 0, "port_id address part of the created controller")
	cmd.Flags().UintVar(&pf, "pf", 0, "physical_function address part of the created controller")
	cmd.Flags().UintVar(&vf, "vf", 0, "virtual_function address part of the created controller")
	cmd.Flags().UintVar(&maxIoQPS, "max-io-qps", 0, "max io queue pairs")

	cobra.CheckErr(cmd.MarkFlagRequired("volume"))
	cobra.CheckErr(cmd.MarkFlagRequired("pf"))
	cobra.CheckErr(cmd.MarkFlagRequired("vf"))

	return cmd
}
