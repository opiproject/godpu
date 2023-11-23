// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package cmd implements the CLI commands
package cmd

import (
	"context"

	"github.com/opiproject/godpu/storage"
	"github.com/spf13/cobra"
)

func newCreateNvmeNamespaceCommand() *cobra.Command {
	id := ""
	subsystem := ""
	volume := ""
	cmd := &cobra.Command{
		Use:     "namespace",
		Aliases: []string{"n"},
		Short:   "Creates nvme namespace",
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

			response, err := client.CreateNvmeNamespace(ctx, id, subsystem, volume)
			cobra.CheckErr(err)

			printResponse(response.Name)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "id for created resource. Assigned by server if omitted.")
	cmd.Flags().StringVar(&subsystem, "subsystem", "", "subsystem name to attach the namespace to")
	cmd.Flags().StringVar(&volume, "volume", "", "volume name to attach as a namespace")

	cobra.CheckErr(cmd.MarkFlagRequired("subsystem"))
	cobra.CheckErr(cmd.MarkFlagRequired("volume"))

	return cmd
}

func newDeleteNvmeNamespaceCommand() *cobra.Command {
	name := ""
	allowMissing := false
	cmd := &cobra.Command{
		Use:     "namespace",
		Aliases: []string{"d"},
		Short:   "Deletes nvme namespace",
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

			err = client.DeleteNvmeNamespace(ctx, name, allowMissing)
			cobra.CheckErr(err)
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "name of deleted namespace")
	cmd.Flags().BoolVar(&allowMissing, "allowMissing", false, "cmd succeeds if attempts to delete a resource that is not present")

	cobra.CheckErr(cmd.MarkFlagRequired("name"))

	return cmd
}
