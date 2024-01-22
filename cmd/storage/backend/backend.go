// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package backend implements the CLI commands for storage backend
package backend

import "github.com/spf13/cobra"

// NewCreateCommand creates a new command to create backend resources
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "backend",
		Aliases: []string{"b"},
		Short:   "Creates backend resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newCreateNvmeCommand())

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

	cmd.AddCommand(newCreateNvmeControllerCommand())
	cmd.AddCommand(newCreateNvmePathCommand())

	return cmd
}

// NewDeleteCommand creates a new command to delete backend resources
func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "backend",
		Aliases: []string{"b"},
		Short:   "Deletes backend resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newDeleteNvmeCommand())

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

	cmd.AddCommand(newDeleteNvmeControllerCommand())
	cmd.AddCommand(newDeleteNvmePathCommand())

	return cmd
}
