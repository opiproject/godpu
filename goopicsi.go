// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package goopicsi implements the go library for OPI to be used in CSI drivers
package goopicsi

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/opiproject/opi-api/storage/v1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ConnectToRemoteAndExpose connects to the remote storage over NVMe/TCP and exposes it as a local NVMe/PCIe device
func ConnectToRemoteAndExpose(addr string) error {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Connect to remote NVMe target from xPU
	c4 := pb.NewNVMfRemoteControllerServiceClient(conn)
	rr0, err := c4.NVMfRemoteControllerConnect(ctx, &pb.NVMfRemoteControllerConnectRequest{Ctrl: &pb.NVMfRemoteController{Id: 1}})
	if err != nil {
		log.Printf("could not connect to Remote NVMf controller: %v", err)
		return err
	}
	log.Printf("Connected: %v", rr0)

	// Expose emulated NVMe device to the Host (Step 1: Subsystem)
	c1 := pb.NewNVMeSubsystemServiceClient(conn)
	rs1, err := c1.NVMeSubsystemCreate(ctx, &pb.NVMeSubsystemCreateRequest{Subsystem: &pb.NVMeSubsystem{Nqn: "OpiMalloc7"}})
	if err != nil {
		log.Printf("could not create NVMe subsystem: %v", err)
		return err
	}
	log.Printf("Added: %v", rs1)
	// Step2: NVMeController
	c2 := pb.NewNVMeControllerServiceClient(conn)
	rc1, err := c2.NVMeControllerCreate(ctx, &pb.NVMeControllerCreateRequest{Controller: &pb.NVMeController{NvmeControllerId: 13}})
	if err != nil {
		log.Printf("could not create NVMe subsystem: %v", err)
		return err
	}
	log.Printf("Added: %v", rc1)

	// NVMeNamespace
	c3 := pb.NewNVMeNamespaceServiceClient(conn)
	rn1, err := c3.NVMeNamespaceCreate(ctx, &pb.NVMeNamespaceCreateRequest{Namespace: &pb.NVMeNamespace{HostNsid: 123}})
	if err != nil {
		log.Printf("could not create NVMe subsystem: %v", err)
		return err
	}
	log.Printf("Added: %v", rn1)
	return nil
}
