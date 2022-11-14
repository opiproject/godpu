// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package goopicsi implements the go library for OPI to be used in CSI drivers
package goopicsi

import (
	"context"
	"flag"
	"log"
	"time"

	pbc "github.com/opiproject/opi-api/common/v1/gen/go"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn    *grpc.ClientConn
	address = "localhost:50051"
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
	c1 := pb.NewFrontendNvmeServiceClient(conn)
	rs1, err := c1.CreateNVMeSubsystem(ctx, &pb.CreateNVMeSubsystemRequest{
		Subsystem: &pb.NVMeSubsystem{
			Spec: &pb.NVMeSubsystemSpec{
				Id:  &pbc.ObjectKey{Value: "controller-test-ss"},
				Nqn: "nqn.2022-09.io.spdk:opi2"}}})
	if err != nil {
		log.Printf("could not create NVMe subsystem: %v", err)
		return err
	}
	log.Printf("Added: %v", rs1)
	// Step2: NVMeController
	rc1, err := c1.CreateNVMeController(ctx, &pb.CreateNVMeControllerRequest{
		Controller: &pb.NVMeController{
			Spec: &pb.NVMeControllerSpec{
				Id:               &pbc.ObjectKey{Value: "controller-test"},
				SubsystemId:      &pbc.ObjectKey{Value: "controller-test-ss"},
				NvmeControllerId: 13}}})
	if err != nil {
		log.Printf("could not create NVMe subsystem: %v", err)
		return err
	}
	log.Printf("Added: %v", rc1)

	// NVMeNamespace
	rn1, err := c1.CreateNVMeNamespace(ctx, &pb.CreateNVMeNamespaceRequest{
		Namespace: &pb.NVMeNamespace{
			Spec: &pb.NVMeNamespaceSpec{
				Id:           &pbc.ObjectKey{Value: "namespace-test"},
				SubsystemId:  &pbc.ObjectKey{Value: "namespace-test-ss"},
				ControllerId: &pbc.ObjectKey{Value: "namespace-test-ctrler"},
				VolumeId:     &pbc.ObjectKey{Value: "Malloc1"},
				HostNsid:     123}}})
	if err != nil {
		log.Printf("could not create NVMe subsystem: %v", err)
		return err
	}
	log.Printf("Added: %v", rn1)
	return nil
}

// NVMeControllerConnect Connects to remote NVMf controller
func NVMeControllerConnect(request *pb.NVMfRemoteController) (*pb.NVMfRemoteControllerConnectResponse, error) {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return nil, err
		}
	}

	client := pb.NewNVMfRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data, err := client.NVMfRemoteControllerGet(ctx, &pb.NVMfRemoteControllerGetRequest{Id: request.Id})
	if err != nil {
		log.Println(err)
	}
	log.Println(data)

	// we will connect if there is no connection established
	if data == nil { // This means we are unable to get a connection with this ID
		response, err := client.NVMfRemoteControllerConnect(ctx, &pb.NVMfRemoteControllerConnectRequest{Ctrl: request})
		if err != nil {
			log.Printf("could not connect to Remote NVMf controller: %v", err)
			return nil, err
		}
		log.Printf("Connected: %v", response)
		return response, nil
	}
	log.Printf("Remote NVMf controller is already connected with SubNQN: %v", data.GetCtrl().Subnqn)
	defer disconnectConnection()
	return &pb.NVMfRemoteControllerConnectResponse{}, nil
}

// NVMeControllerList lists all the connections to the remote NVMf controller
func NVMeControllerList() error {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return err
		}
	}

	client := pb.NewNVMfRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.NVMfRemoteControllerList(ctx, &pb.NVMfRemoteControllerListRequest{})
	if err != nil {
		log.Printf("could not list the connections to Remote NVMf controller: %v", err)
		return err
	}
	log.Printf("Connections: %v", response)
	return nil
}

// NVMeControllerGet lists the connection to the remote NVMf controller corresponding to the given ID
func NVMeControllerGet(id int64) error {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return err
		}
	}

	client := pb.NewNVMfRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.NVMfRemoteControllerGet(ctx, &pb.NVMfRemoteControllerGetRequest{Id: id})
	if err != nil {
		log.Printf("could not list the connection to Remote NVMf controller corresponding to the given ID: %v", err)
		return err
	}
	log.Printf("Connection corresponding to the given ID: %v", response)
	return nil
}

// NVMeControllerDisconnect disconnects remote NVMf controller connection
func NVMeControllerDisconnect(request *pb.NVMfRemoteControllerDisconnectRequest) (*pb.NVMfRemoteControllerDisconnectResponse, error) {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return nil, err
		}
	}

	client := pb.NewNVMfRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data, err := client.NVMfRemoteControllerGet(ctx, &pb.NVMfRemoteControllerGetRequest{Id: request.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(data)

	// we will disconnect if there is a connection
	if data != nil {
		response, err := client.NVMfRemoteControllerDisconnect(ctx, &pb.NVMfRemoteControllerDisconnectRequest{Id: request.Id})
		if err != nil {
			log.Printf("could not disconnect Remote NVMf controller: %v", err)
			return nil, err
		}
		log.Printf("disconnected: %v", response)
		return response, nil
	}
	log.Printf("Remote NVMf controller disconnected successfully: %v", data.GetCtrl().Subnqn)
	defer disconnectConnection()
	return &pb.NVMfRemoteControllerDisconnectResponse{}, nil
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
	} else {
		log.Println("GRPC connection closed successfully")
	}
}
