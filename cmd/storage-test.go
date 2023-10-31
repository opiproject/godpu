// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (C) 2023 Intel Corporation

// Package cmd implements the CLI commands
package cmd

import (
	"context"
	"log"
	"time"

	"github.com/opiproject/godpu/grpc"
	"github.com/opiproject/godpu/storage"
	"github.com/spf13/cobra"
)

const addrCmdLineArg = "addr"

type storagePartition string

const (
	storagePartitionFrontend  storagePartition = "frontend"
	storagePartitionBackend   storagePartition = "backend"
	storagePartitionMiddleend storagePartition = "middleend"
)

var allStoragePartitions = []storagePartition{
	storagePartitionFrontend,
	storagePartitionBackend,
	storagePartitionMiddleend,
}

// NewStorageCommand tests the storage functionality
func NewStorageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "storage",
		Aliases: []string{"g"},
		Short:   "Tests storage functionality",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			if err != nil {
				log.Fatalf("[ERROR] %s", err.Error())
			}
		},
	}

	cmd.AddCommand(newStorageTestCommand())

	return cmd
}

func newStorageTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "test",
		Aliases: []string{"s"},
		Short:   "Test storage functionality",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			runTests(
				c,
				allStoragePartitions,
				storage.AllFrontendPartitions,
			)
		},
	}

	cmd.AddCommand(newStorageTestFrontendCommand())
	cmd.AddCommand(newStorageTestBackendCommand())
	cmd.AddCommand(newStorageTestMiddleendCommand())

	flags := cmd.PersistentFlags()
	flags.String(addrCmdLineArg, "localhost:50151", "address of OPI gRPC server")
	return cmd
}

func newStorageTestFrontendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   string(storagePartitionFrontend),
		Short: "Tests storage frontend API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionFrontend},
				storage.AllFrontendPartitions,
			)
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "nvme",
		Short: "Tests storage frontend nvme API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionFrontend},
				[]storage.FrontendPartition{storage.FrontendPartitionNvme},
			)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "virtio-blk",
		Short: "Tests storage frontend virtio-blk API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionFrontend},
				[]storage.FrontendPartition{storage.FrontendPartitionVirtioBlk},
			)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "scsi",
		Short: "Tests storage frontend scsi API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionFrontend},
				[]storage.FrontendPartition{storage.FrontendPartitionScsi},
			)
		},
	})

	return cmd
}

func newStorageTestBackendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   string(storagePartitionBackend),
		Short: "Tests storage backend API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionBackend},
				nil,
			)
		},
	}

	return cmd
}

func newStorageTestMiddleendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   string(storagePartitionMiddleend),
		Short: "Tests storage middleend API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionMiddleend},
				nil,
			)
		},
	}

	return cmd
}

func runTests(
	cmd *cobra.Command,
	partitions []storagePartition,
	frontendPartitions []storage.FrontendPartition,
) {
	addr, err := cmd.Flags().GetString(addrCmdLineArg)
	if err != nil {
		log.Fatalf("error getting %v argument: %v", addrCmdLineArg, err)
	}

	// Set up a connection to the server.
	client, err := grpc.New(addr)
	if err != nil {
		log.Fatalf("error creating new client: %v", err)
	}

	// Contact the server and print out its response.
	conn, closer, err := client.NewConn()
	if err != nil {
		log.Fatalf("error creating gRPC connection: %v", err)
	}
	defer closer()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, partition := range partitions {
		log.Printf("==============================================================================")
		log.Printf("Test %v", partition)
		log.Printf("==============================================================================")

		var err error
		switch partition {
		case storagePartitionFrontend:
			err = storage.DoFrontend(ctx, conn, frontendPartitions)
		case storagePartitionBackend:
			err = storage.DoBackend(ctx, conn)
		case storagePartitionMiddleend:
			err = storage.DoMiddleend(ctx, conn)
		default:
			log.Panicf("Unknown storage partition: %v", partition)
		}

		if err != nil {
			log.Panicf("%v tests failed with error: %v", partition, err)
		}
	}
}
