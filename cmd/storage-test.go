// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package cmd implements the CLI commands
package cmd

import (
	"context"
	"log"
	"time"

	"github.com/opiproject/godpu/storage"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewStorageTestCommand returns the storage tests command
func NewStorageTestCommand() *cobra.Command {
	var (
		addr string
	)
	cmd := &cobra.Command{
		Use:     "storagetest",
		Aliases: []string{"s"},
		Short:   "Test storage functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// Set up a connection to the server.
			conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer func(conn *grpc.ClientConn) {
				err := conn.Close()
				if err != nil {
					log.Fatalf("did not close connection: %v", err)
				}
			}(conn)

			// Contact the server and print out its response.
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			log.Printf("==============================================================================")
			log.Printf("Test frontend")
			log.Printf("==============================================================================")
			err = storage.DoFrontend(ctx, conn)
			if err != nil {
				log.Panicf("DoFrontend tests failed with error: %v", err)
			}

			log.Printf("==============================================================================")
			log.Printf("Test backend")
			log.Printf("==============================================================================")
			err = storage.DoBackend(ctx, conn)
			if err != nil {
				log.Panicf("DoFrontend tests failed with error: %v", err)
			}

			log.Printf("==============================================================================")
			log.Printf("Test middleend")
			log.Printf("==============================================================================")
			err = storage.DoMiddleend(ctx, conn)
			if err != nil {
				log.Panicf("DoFrontend tests failed with error: %v", err)
			}
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&addr, "addr", "localhost:50151", "address or OPI gRPC server")
	return cmd
}
