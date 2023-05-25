// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

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

// NvmeConnection defines remote NVMf connection
type NvmeConnection struct {
	id     string
	subnqn string
	traddr string
}

// NvmeControllerConnect Connects to remote NVMf controller
func NvmeControllerConnect(id string, trAddr string, subnqn string, trSvcID int64, hostnqn string) error {
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
		request := &pb.CreateNVMfRemoteControllerRequest{NvMfRemoteControllerId: id, NvMfRemoteController: &pb.NVMfRemoteController{
			Name:    id,
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

// NvmeControllerList lists all the connections to the remote NVMf controller
func NvmeControllerList() ([]NvmeConnection, error) {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return []NvmeConnection{}, err
		}
	}

	client := pb.NewNVMfRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.ListNVMfRemoteControllers(ctx, &pb.ListNVMfRemoteControllersRequest{})
	if err != nil {
		log.Printf("could not list the connections to Remote NVMf controller: %v", err)
		return []NvmeConnection{}, err
	}
	nvmeConnections := make([]NvmeConnection, 0)
	for _, connection := range response.NvMfRemoteControllers {
		nvmeConnections = append(nvmeConnections, NvmeConnection{
			id:     "",
			subnqn: connection.Subnqn,
			traddr: "",
		})
	}
	return nvmeConnections, nil
}

// NvmeControllerGet lists the connection to the remote NVMf controller corresponding to the given ID
func NvmeControllerGet(id string) (string, error) {
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

// NvmeControllerDisconnect disconnects remote NVMf controller connection
func NvmeControllerDisconnect(id string) error {
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

// ExposeRemoteNvme creates a new Nvme Subsystem and Nvme controller. Default value of MaxNamespaces is 32 incase the parameter is not assigned any value
func ExposeRemoteNvme(subsystemNQN string, maxNamespaces int64) (string, string, error) {
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
	data1, err := client.GetNvmeSubsystem(ctx, &pb.GetNvmeSubsystemRequest{Name: subsystemID})
	if err != nil {
		log.Printf("No existing Nvme Subsystem found with subsystemID: %v", subsystemID)
	}

	if data1 == nil {
		response1, err := client.CreateNvmeSubsystem(ctx, &pb.CreateNvmeSubsystemRequest{
			NvmeSubsystem: &pb.NvmeSubsystem{
				Spec: &pb.NvmeSubsystemSpec{
					Name:          subsystemID,
					Nqn:           subsystemNQN,
					MaxNamespaces: maxNamespaces,
				},
			},
		})
		if err != nil {
			log.Println(err)
			return "", "", err
		}
		log.Printf("Nvme Subsytem created: %v", response1)
	} else {
		log.Printf("Nvme Subsystem is already present with the subsytemID: %v", subsystemID)
	}

	controllerID := uuid.New().String()
	data2, err := client.GetNvmeController(ctx, &pb.GetNvmeControllerRequest{Name: controllerID})
	if err != nil {
		log.Printf("No existing Nvme Controller found with controllerID: %v", controllerID)
	}

	// Default value of MaxNamespaces is 32 incase the parameter is not assigned any value
	if data2 == nil {
		response2, err := client.CreateNvmeController(ctx, &pb.CreateNvmeControllerRequest{
			NvmeController: &pb.NvmeController{
				Spec: &pb.NvmeControllerSpec{
					Name:          controllerID,
					SubsystemId:   &pbc.ObjectKey{Value: subsystemID},
					MaxNamespaces: int32(maxNamespaces),
				},
			},
		})
		if err != nil {
			log.Println(err)
			return subsystemID, "", err
		}
		log.Printf("Nvme Controller created: %v", response2)
		return subsystemID, controllerID, nil
	}
	log.Printf("Nvme Controller is already present with the controllerID: %v", controllerID)
	return subsystemID, controllerID, nil
}

// CreateNvmeNamespace Creates a new Nvme namespace
func CreateNvmeNamespace(id string, subSystemID string, nguid string, hostID int32) (string, error) {
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
			volumeID = data.Name
		}
	}
	if volumeID == "" {
		return "", errors.New("volume ID not found")
	}

	client2 := pb.NewFrontendNvmeServiceClient(conn)
	resp, err := client2.CreateNvmeNamespace(ctx, &pb.CreateNvmeNamespaceRequest{
		NvmeNamespace: &pb.NvmeNamespace{
			Spec: &pb.NvmeNamespaceSpec{
				Name:        id,
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
	return resp.Spec.Name, nil
}

// DeleteNvmeNamespace deletes the Nvme namespace
func DeleteNvmeNamespace(id string) error {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return err
		}
	}

	client := pb.NewFrontendNvmeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.DeleteNvmeNamespace(ctx, &pb.DeleteNvmeNamespaceRequest{Name: id})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(resp)
	return nil
}

// GenerateHostNQN generates a new hostNQN
func GenerateHostNQN() string {
	// Sample of Nvme Qualified Name in UUID-based format - nqn.2014-08.org.nvmexpress:uuid:a11a1111-11a1-111a-a111-1a111aaa1a11
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
	}
	log.Println("GRPC connection closed successfully")
}
