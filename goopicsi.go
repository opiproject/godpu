// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package goopicsi implements the go library for OPI to be used in CSI drivers
package goopicsi

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"

	pbc "github.com/opiproject/opi-api/common/v1/gen/go"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn    *grpc.ClientConn
	address = "localhost:50051"
)

// NVMeConnection defines remote NVMf connection
type NVMeConnection struct {
	id     string
	subnqn string
	traddr string
}

// NVMeControllerConnect Connects to remote NVMf controller
func NVMeControllerConnect(id string, trAddr string, subnqn string, trSvcID int64, hostnqn string) error {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return err
		}
	}

	client := pb.NewNVMfRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data, err := client.GetNVMfRemoteController(ctx, &pb.GetNVMfRemoteControllerRequest{Name: id})
	if err != nil {
		log.Println(err)
	}
	log.Println(data)

	// we will connect if there is no connection established
	if data == nil { // This means we are unable to get a connection with this ID
		request := &pb.CreateNVMfRemoteControllerRequest{NvMfRemoteController: &pb.NVMfRemoteController{
			Id:      &pbc.ObjectKey{Value: id},
			Traddr:  trAddr,
			Subnqn:  subnqn,
			Trsvcid: trSvcID,
			Hostnqn: hostnqn,
			Trtype:  pb.NvmeTransportType_NVME_TRANSPORT_TCP,
			Adrfam:  pb.NvmeAddressFamily_NVMF_ADRFAM_IPV4,
		}}
		response, err := client.CreateNVMfRemoteController(ctx, request)
		if err != nil {
			log.Printf("could not connect to Remote NVMf controller: %v", err)
			return err
		}
		log.Printf("Connected: %v", response)
		return nil
	}
	log.Printf("Remote NVMf controller is already connected with SubNQN: %v", data.Subnqn)

	return nil
}

// NVMeControllerList lists all the connections to the remote NVMf controller
func NVMeControllerList() ([]NVMeConnection, error) {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return []NVMeConnection{}, err
		}
	}

	client := pb.NewNVMfRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.ListNVMfRemoteControllers(ctx, &pb.ListNVMfRemoteControllersRequest{})
	if err != nil {
		log.Printf("could not list the connections to Remote NVMf controller: %v", err)
		return []NVMeConnection{}, err
	}
	nvmeConnections := make([]NVMeConnection, 0)
	for _, connection := range response.NvMfRemoteControllers {
		nvmeConnections = append(nvmeConnections, NVMeConnection{
			id:     "",
			subnqn: connection.Subnqn,
			traddr: "",
		})
	}
	return nvmeConnections, nil
}

// NVMeControllerGet lists the connection to the remote NVMf controller corresponding to the given ID
func NVMeControllerGet(id string) (string, error) {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return "", err
		}
	}

	client := pb.NewNVMfRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetNVMfRemoteController(ctx, &pb.GetNVMfRemoteControllerRequest{Name: id})
	if err != nil {
		log.Printf("could not list the connection to Remote NVMf controller corresponding to the given ID: %v", err)
		return "", err
	}
	return response.Subnqn, nil
}

// NVMeControllerDisconnect disconnects remote NVMf controller connection
func NVMeControllerDisconnect(id string) error {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return err
		}
	}

	client := pb.NewNVMfRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data, err := client.GetNVMfRemoteController(ctx, &pb.GetNVMfRemoteControllerRequest{Name: id})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(data)

	// we will disconnect if there is a connection
	if data != nil {
		response, err := client.DeleteNVMfRemoteController(ctx, &pb.DeleteNVMfRemoteControllerRequest{Name: id})
		if err != nil {
			log.Printf("could not disconnect Remote NVMf controller: %v", err)
			return err
		}
		log.Printf("disconnected: %v", response)
		return nil
	}
	log.Printf("Remote NVMf controller disconnected successfully: %v", data.Subnqn)
	defer disconnectConnection()
	return nil
}

// ExposeRemoteNVMe creates a new NVMe Subsystem and NVMe controller
func ExposeRemoteNVMe(subsystemNQN string, maxNamespaces int64) (string, string, error) {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return "", "", err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewFrontendNvmeServiceClient(conn)
	subsystemID := uuid.New().String()
	data1, err := client.GetNVMeSubsystem(ctx, &pb.GetNVMeSubsystemRequest{Name: subsystemID})
	if err != nil {
		log.Printf("No existing NVMe Subsystem found with subsystemID: %v", subsystemID)
	}

	if data1 == nil {
		response1, err := client.CreateNVMeSubsystem(ctx, &pb.CreateNVMeSubsystemRequest{
			NvMeSubsystem: &pb.NVMeSubsystem{
				Spec: &pb.NVMeSubsystemSpec{
					Id:            &pbc.ObjectKey{Value: subsystemID},
					Nqn:           subsystemNQN,
					MaxNamespaces: maxNamespaces,
				},
			},
		})
		if err != nil {
			log.Println(err)
			return "", "", err
		}
		log.Printf("NVMe Subsytem created: %v", response1)
	} else {
		log.Printf("NVMe Subsystem is already present with the subsytemID: %v", subsystemID)
	}

	controllerID := uuid.New().String()
	data2, err := client.GetNVMeController(ctx, &pb.GetNVMeControllerRequest{Name: controllerID})
	if err != nil {
		log.Printf("No existing NVMe Controller found with controllerID: %v", controllerID)
	}

	// Default value of MaxNamespaces is 32 incase the parameter is not assigned any value
	if data2 == nil {
		response2, err := client.CreateNVMeController(ctx, &pb.CreateNVMeControllerRequest{
			NvMeController: &pb.NVMeController{
				Spec: &pb.NVMeControllerSpec{
					Id:            &pbc.ObjectKey{Value: controllerID},
					SubsystemId:   &pbc.ObjectKey{Value: subsystemID},
					MaxNamespaces: int32(maxNamespaces),
				},
			},
		})
		if err != nil {
			log.Println(err)
			return subsystemID, "", err
		}
		log.Printf("NVMe Controller created: %v", response2)
		return subsystemID, controllerID, nil
	}
	log.Printf("NVMe Controller is already present with the controllerID: %v", controllerID)
	return subsystemID, controllerID, nil
}

// CreateNVMeNamespace Creates a new NVMe namespace
func CreateNVMeNamespace(id string, subSystemID string, nguid string, hostID int32) (string, error) {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return "", err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client1 := pb.NewNullDebugServiceClient(conn)
	response, err := client1.ListNullDebugs(ctx, &pb.ListNullDebugsRequest{})

	if err != nil {
		log.Println(err)
		return "", err
	}

	volumeData := response.NullDebugs
	volumeID := ""
	for _, data := range volumeData {
		uuid := strings.ReplaceAll(data.Uuid.Value, "-", "")
		if uuid == nguid {
			volumeID = data.Handle.Value
		}
	}
	if volumeID == "" {
		return "", errors.New("volume ID not found")
	}

	client2 := pb.NewFrontendNvmeServiceClient(conn)
	resp, err := client2.CreateNVMeNamespace(ctx, &pb.CreateNVMeNamespaceRequest{
		NvMeNamespace: &pb.NVMeNamespace{
			Spec: &pb.NVMeNamespaceSpec{
				Id:          &pbc.ObjectKey{Value: id},
				SubsystemId: &pbc.ObjectKey{Value: subSystemID},
				VolumeId:    &pbc.ObjectKey{Value: volumeID},
				HostNsid:    hostID,
			},
		},
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(resp)
	return resp.Spec.Id.Value, nil
}

// DeleteNVMeNamespace deletes the NVMe namespace
func DeleteNVMeNamespace(id string) error {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return err
		}
	}

	client := pb.NewFrontendNvmeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.DeleteNVMeNamespace(ctx, &pb.DeleteNVMeNamespaceRequest{Name: id})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(resp)
	return nil
}

// GenerateHostNQN generates a new hostNQN
func GenerateHostNQN() string {
	// Sample of NVMe Qualified Name in UUID-based format - nqn.2014-08.org.nvmexpress:uuid:a11a1111-11a1-111a-a111-1a111aaa1a11
	nqnConst := "nqn.2014-08.org.nvmexpress:uuid:"
	nqnUUID := uuid.New().String()

	hostNQN := fmt.Sprintf("%s%s", nqnConst, nqnUUID)
	return hostNQN
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
