// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package cmd implements the CLI commands
package cmd

import (
	"context"
	"net"

	"github.com/opiproject/godpu/storage"
	"github.com/spf13/cobra"
)

func newCreateNvmeControllerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "controller",
		Aliases: []string{"c"},
		Short:   "Creates nvme controller",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newCreateNvmeControllerTCPCommand())

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
		Run: func(c *cobra.Command, args []string) {
			addr, err := c.Flags().GetString(addrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(timeoutCmdLineArg)
			cobra.CheckErr(err)

			client, err := storage.New(addr)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			response, err := client.CreateNvmeTCPController(ctx, id, subsystem, ip, port)
			cobra.CheckErr(err)

			printResponse(response.Name)
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

func newDeleteNvmeControllerCommand() *cobra.Command {
	name := ""
	allowMissing := false
	cmd := &cobra.Command{
		Use:     "controller",
		Aliases: []string{"c"},
		Short:   "Deletes nvme controller",
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

			err = client.DeleteNvmeController(ctx, name, allowMissing)
			cobra.CheckErr(err)
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "name of deleted controller")
	cmd.Flags().BoolVar(&allowMissing, "allowMissing", false, "cmd succeeds if attempts to delete a resource that is not present")

	cobra.CheckErr(cmd.MarkFlagRequired("name"))

	return cmd
}
