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
			)
		},
	}

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
			)
		},
	}

	return cmd
}

func runTests(
	cmd *cobra.Command,
	partitions []storagePartition,
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
			err = storage.DoFrontend(ctx, conn)
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
