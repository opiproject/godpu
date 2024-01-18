// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023-2024 Intel Corporation

// Package frontend implements the CLI commands for storage frontend
package frontend

import "github.com/spf13/cobra"

func NewCreateNvmeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nvme",
		Aliases: []string{"n"},
		Short:   "Creates nvme resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newCreateNvmeSubsystemCommand())
	cmd.AddCommand(newCreateNvmeNamespaceCommand())
	cmd.AddCommand(newCreateNvmeControllerCommand())

	return cmd
}

func NewCreateVirtioCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "virtio",
		Aliases: []string{"v"},
		Short:   "Creates virtio resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newCreateVirtioBlkCommand())

	return cmd
}

func NewDeleteNvmeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nvme",
		Aliases: []string{"n"},
		Short:   "Deletes nvme resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newDeleteNvmeSubsystemCommand())
	cmd.AddCommand(newDeleteNvmeNamespaceCommand())
	cmd.AddCommand(newDeleteNvmeControllerCommand())

	return cmd
}

func NewDeleteVirtioCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "virtio",
		Aliases: []string{"v"},
		Short:   "Deletes virtio resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newDeleteVirtioBlkCommand())

	return cmd
}
