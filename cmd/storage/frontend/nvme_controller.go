// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023-2024 Intel Corporation

// Package frontend implements the CLI commands for storage frontend
package frontend

import (
	"context"
	"net"

	"github.com/opiproject/godpu/cmd/storage/common"
	frontendclient "github.com/opiproject/godpu/storage/frontend"
	"github.com/spf13/cobra"
)

func newCreateNvmeControllerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "controller",
		Aliases: []string{"c"},
		Short:   "Creates nvme controller",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newCreateNvmeControllerTCPCommand())
	cmd.AddCommand(newCreateNvmeControllerPcieCommand())

	return cmd
}

func newCreateNvmeControllerTCPCommand() *cobra.Command {
	id := ""
	subsystem := ""
	var ip net.IP
	var port uint16
	cmd := &cobra.Command{
		Use:     "tcp",
		Aliases: []string{"t"},
		Short:   "Creates nvme TCP controller",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(common.TimeoutCmdLineArg)
			cobra.CheckErr(err)

			client, err := frontendclient.New(addr)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			response, err := client.CreateNvmeTCPController(ctx, id, subsystem, ip, port)
			cobra.CheckErr(err)

			common.PrintResponse(response.Name)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "id for created resource. Assigned by server if omitted.")
	cmd.Flags().StringVar(&subsystem, "subsystem", "", "subsystem name to attach the controller to")
	cmd.Flags().IPVar(&ip, "ip", nil, "ip address of the created controller")
	cmd.Flags().Uint16Var(&port, "port", 0, "port of the created controller")

	cobra.CheckErr(cmd.MarkFlagRequired("subsystem"))
	cobra.CheckErr(cmd.MarkFlagRequired("ip"))
	cobra.CheckErr(cmd.MarkFlagRequired("port"))

	return cmd
}

func newCreateNvmeControllerPcieCommand() *cobra.Command {
	id := ""
	subsystem := ""
	var port uint
	var pf uint
	var vf uint
	cmd := &cobra.Command{
		Use:     "pcie",
		Aliases: []string{"p"},
		Short:   "Creates nvme PCIe controller",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(common.TimeoutCmdLineArg)
			cobra.CheckErr(err)

			client, err := frontendclient.New(addr)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			response, err := client.CreateNvmePcieController(ctx, id, subsystem, port, pf, vf)
			cobra.CheckErr(err)

			common.PrintResponse(response.Name)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "id for created resource. Assigned by server if omitted.")
	cmd.Flags().StringVar(&subsystem, "subsystem", "", "subsystem name to attach the controller to")
	cmd.Flags().UintVar(&port, "port", 0, "port_id address part of the created controller")
	cmd.Flags().UintVar(&pf, "pf", 0, "physical_function address part of the created controller")
	cmd.Flags().UintVar(&vf, "vf", 0, "virtual_function address part of the created controller")

	cobra.CheckErr(cmd.MarkFlagRequired("subsystem"))
	cobra.CheckErr(cmd.MarkFlagRequired("pf"))
	cobra.CheckErr(cmd.MarkFlagRequired("vf"))

	return cmd
}

func newDeleteNvmeControllerCommand() *cobra.Command {
	name := ""
	allowMissing := false
	cmd := &cobra.Command{
		Use:     "controller",
		Aliases: []string{"c"},
		Short:   "Deletes nvme controller",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(common.TimeoutCmdLineArg)
			cobra.CheckErr(err)

			client, err := frontendclient.New(addr)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			err = client.DeleteNvmeController(ctx, name, allowMissing)
			cobra.CheckErr(err)
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "name of deleted controller")
	cmd.Flags().BoolVar(&allowMissing, "allowMissing", false, "cmd succeeds if attempts to delete a resource that is not present")

	cobra.CheckErr(cmd.MarkFlagRequired("name"))

	return cmd
}
