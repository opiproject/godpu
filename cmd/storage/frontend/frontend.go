// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023-2024 Intel Corporation

// Package frontend implements the CLI commands for storage frontend
package frontend

import "github.com/spf13/cobra"

// NewCreateCommand creates a new command to create frontend resources
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "frontend",
		Aliases: []string{"f"},
		Short:   "Creates frontend resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newCreateNvmeCommand())
	cmd.AddCommand(newCreateVirtioCommand())

	return cmd
}

func newCreateNvmeCommand() *cobra.Command {
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

func newCreateVirtioCommand() *cobra.Command {
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

// NewDeleteCommand creates a new command to delete frontend resources
func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "frontend",
		Aliases: []string{"f"},
		Short:   "Deletes frontend resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newDeleteNvmeCommand())
	cmd.AddCommand(newDeleteVirtioCommand())

	return cmd
}

func newDeleteNvmeCommand() *cobra.Command {
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

func newDeleteVirtioCommand() *cobra.Command {
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
