// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package inventory implements the go library for OPI to be used to query inventory
package inventory

import (
	"context"
	"log"
	"time"

	pb "github.com/opiproject/opi-api/common/v1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn    *grpc.ClientConn
	address = "localhost:50051"
)

// Get returns inventory information from DPUs
func Get() error {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return err
		}
	}

	client := pb.NewInventorySvcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data, err := client.InventoryGet(ctx, &pb.InventoryGetRequest{})
	if err != nil {
		log.Println(err)
	}
	log.Println(data)
	defer disconnectConnection()
	return nil
}

func dialConnection() error {
	var err error
	conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect: %v", err)
		return err
	}
	return nil
}

func disconnectConnection() {
	err := conn.Close()
	if err != nil {
		log.Fatalf("Failed to close connection: %v", err)
	}
	log.Println("GRPC connection closed successfully")
}
