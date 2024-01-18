// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package frontend implements the CLI commands for storage frontend
package frontend

import (
	"context"

	"github.com/opiproject/godpu/cmd/storage/common"
	frontendclient "github.com/opiproject/godpu/storage/frontend"
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
			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(common.TimeoutCmdLineArg)
			cobra.CheckErr(err)

			client, err := frontendclient.New(addr)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			response, err := client.CreateVirtioBlk(ctx, id, volume, port, pf, vf, maxIoQPS)
			cobra.CheckErr(err)

			common.PrintResponse(response.Name)
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

func newDeleteVirtioBlkCommand() *cobra.Command {
	name := ""
	allowMissing := false
	cmd := &cobra.Command{
		Use:     "blk",
		Aliases: []string{"b"},
		Short:   "Deletes virtio-blk controller",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(common.TimeoutCmdLineArg)
			cobra.CheckErr(err)

			client, err := frontendclient.New(addr)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			err = client.DeleteVirtioBlk(ctx, name, allowMissing)
			cobra.CheckErr(err)
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "name of deleted virtio-blk controller")
	cmd.Flags().BoolVar(&allowMissing, "allowMissing", false, "cmd succeeds if attempts to delete a resource that is not present")

	cobra.CheckErr(cmd.MarkFlagRequired("name"))

	return cmd
}
