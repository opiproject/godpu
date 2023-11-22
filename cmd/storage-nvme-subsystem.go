// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package cmd implements the CLI commands
package cmd

import (
	"context"

	"github.com/opiproject/godpu/storage"
	"github.com/spf13/cobra"
)

func newCreateNvmeSubsystemCommand() *cobra.Command {
	id := ""
	nqn := ""
	hostnqn := ""
	cmd := &cobra.Command{
		Use:     "subsystem",
		Aliases: []string{"s"},
		Short:   "Creates nvme subsystem",
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

			response, err := client.CreateNvmeSubsystem(ctx, id, nqn, hostnqn)
			cobra.CheckErr(err)

			printResponse(response.Name)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "id for created resource. Assigned by server if omitted.")
	cmd.Flags().StringVar(&nqn, "nqn", "", "nqn for created subsystem")
	cmd.Flags().StringVar(&hostnqn, "hostnqn", "", "hostnqn for created subsystem")

	cobra.CheckErr(cmd.MarkFlagRequired("nqn"))

	return cmd
}

func newDeleteNvmeSubsystemCommand() *cobra.Command {
	name := ""
	allowMissing := false
	cmd := &cobra.Command{
		Use:     "subsystem",
		Aliases: []string{"s"},
		Short:   "Deletes nvme subsystem",
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

			err = client.DeleteNvmeSubsystem(ctx, name, allowMissing)
			cobra.CheckErr(err)
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "name of deleted subsystem")
	cmd.Flags().BoolVar(&allowMissing, "allowMissing", false, "cmd succeeds if attempts to delete a resource that is not present")

	cobra.CheckErr(cmd.MarkFlagRequired("name"))

	return cmd
}
