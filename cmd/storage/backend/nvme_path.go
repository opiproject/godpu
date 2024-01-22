// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package backend implements the CLI commands for storage backend
package backend

import (
	"context"
	"net"

	"github.com/opiproject/godpu/cmd/storage/common"
	backendclient "github.com/opiproject/godpu/storage/backend"
	"github.com/spf13/cobra"
)

func newCreateNvmePathCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "path",
		Aliases: []string{"p"},
		Short:   "Creates nvme path to an external nvme device",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newCreateNvmePathTCPCommand())

	return cmd
}

func newCreateNvmePathTCPCommand() *cobra.Command {
	id := ""
	nqn := ""
	hostnqn := ""
	controller := ""
	var ip net.IP
	var port uint16
	cmd := &cobra.Command{
		Use:     "tcp",
		Aliases: []string{"t"},
		Short:   "Creates nvme path to a remote nvme TCP controller",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(common.TimeoutCmdLineArg)
			cobra.CheckErr(err)

			client, err := backendclient.New(addr)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			response, err := client.CreateNvmeTCPPath(ctx, id, controller, ip, port, nqn, hostnqn)
			cobra.CheckErr(err)

			common.PrintResponse(response.Name)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "id for created resource. Assigned by server if omitted.")
	cmd.Flags().StringVar(&controller, "controller", "", "backend controller name for this path")
	cmd.Flags().IPVar(&ip, "ip", nil, "ip address of the path to connect to.")
	cmd.Flags().Uint16Var(&port, "port", 0, "port of the path to connect to.")
	cmd.Flags().StringVar(&nqn, "nqn", "", "nqn of the target subsystem.")
	cmd.Flags().StringVar(&hostnqn, "hostnqn", "", "host nqn")

	cobra.CheckErr(cmd.MarkFlagRequired("controller"))
	cobra.CheckErr(cmd.MarkFlagRequired("ip"))
	cobra.CheckErr(cmd.MarkFlagRequired("port"))
	cobra.CheckErr(cmd.MarkFlagRequired("nqn"))

	return cmd
}

func newDeleteNvmePathCommand() *cobra.Command {
	name := ""
	allowMissing := false
	cmd := &cobra.Command{
		Use:     "path",
		Aliases: []string{"p"},
		Short:   "Deletes nvme path to an external nvme device",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(common.TimeoutCmdLineArg)
			cobra.CheckErr(err)

			client, err := backendclient.New(addr)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			err = client.DeleteNvmePath(ctx, name, allowMissing)
			cobra.CheckErr(err)
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "name of deleted nvme path")
	cmd.Flags().BoolVar(&allowMissing, "allowMissing", false, "cmd succeeds if attempts to delete a resource that is not present")

	cobra.CheckErr(cmd.MarkFlagRequired("name"))

	return cmd
}
