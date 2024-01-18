// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (C) 2023 Intel Corporation

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

	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn    *grpc.ClientConn
	address = "localhost:50051"
)

// NvmeConnection defines remote Nvme connection
type NvmeConnection struct {
	id     string
	subnqn string
	traddr string
}

// NvmeControllerConnect Connects to remote Nvme controller
func NvmeControllerConnect(id string, trAddr string, subnqn string, trSvcID int64, hostnqn string) error {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return err
		}
	}

	client := pb.NewNvmeRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data, err := client.GetNvmeRemoteController(ctx, &pb.GetNvmeRemoteControllerRequest{Name: id})
	if err != nil {
		log.Println(err)
	}
	log.Println(data)

	// we will connect if there is no connection established
	if data == nil { // This means we are unable to get a connection with this ID
		request := &pb.CreateNvmeRemoteControllerRequest{NvmeRemoteControllerId: id, NvmeRemoteController: &pb.NvmeRemoteController{
			Name:      id,
			Multipath: pb.NvmeMultipath_NVME_MULTIPATH_DISABLE,
		}}
		response, err := client.CreateNvmeRemoteController(ctx, request)
		if err != nil {
			log.Printf("could not connect to Remote Nvme controller: %v", err)
			return err
		}
		log.Printf("Connected: %v", response)

		pathResponse, err := client.CreateNvmePath(ctx, &pb.CreateNvmePathRequest{
			Parent:     response.Name,
			NvmePathId: nvmeControllerToPathResourceID(id),
			NvmePath: &pb.NvmePath{
				Traddr: trAddr,
				Trtype: pb.NvmeTransportType_NVME_TRANSPORT_TYPE_TCP,
				Fabrics: &pb.FabricsPath{
					Subnqn:  subnqn,
					Trsvcid: trSvcID,
					Hostnqn: hostnqn,
					Adrfam:  pb.NvmeAddressFamily_NVME_ADDRESS_FAMILY_IPV4,
				},
			},
		})
		if err != nil {
			log.Printf("could not connect to Remote Nvme path: %v", err)
			_, _ = client.DeleteNvmeRemoteController(ctx, &pb.DeleteNvmeRemoteControllerRequest{
				Name: response.Name,
			})
			return err
		}
		log.Printf("Connected: %v", pathResponse)

		return nil
	}
	log.Printf("Remote Nvme controller is already connected")

	return nil
}

// NvmeControllerList lists all the connections to the remote Nvme controller
func NvmeControllerList() ([]NvmeConnection, error) {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return []NvmeConnection{}, err
		}
	}

	client := pb.NewNvmeRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.ListNvmeRemoteControllers(ctx, &pb.ListNvmeRemoteControllersRequest{})
	if err != nil {
		log.Printf("could not list the connections to Remote Nvme controller: %v", err)
		return []NvmeConnection{}, err
	}
	nvmeConnections := make([]NvmeConnection, 0)
	for range response.NvmeRemoteControllers {
		nvmeConnections = append(nvmeConnections, NvmeConnection{
			id: "",
			// TODO: fetch Nvme paths to fill when OPI API is extended with List/Get calls
			subnqn: "",
			traddr: "",
		})
	}
	return nvmeConnections, nil
}

// NvmeControllerGet lists the connection to the remote Nvme controller corresponding to the given ID
func NvmeControllerGet(id string) (string, error) {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return "", err
		}
	}

	client := pb.NewNvmeRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.GetNvmeRemoteController(ctx, &pb.GetNvmeRemoteControllerRequest{Name: id})
	if err != nil {
		log.Printf("could not list the connection to Remote Nvme controller corresponding to the given ID: %v", err)
		return "", err
	}
	// TODO: fetch nqn in Nvme path when OPI API is extended with List/Get calls
	return "", err
}

// NvmeControllerDisconnect disconnects remote Nvme controller connection
func NvmeControllerDisconnect(id string) error {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return err
		}
	}

	client := pb.NewNvmeRemoteControllerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data, err := client.GetNvmeRemoteController(ctx, &pb.GetNvmeRemoteControllerRequest{Name: id})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(data)

	// we will disconnect if there is a connection
	if data != nil {
		_, err := client.DeleteNvmePath(ctx, &pb.DeleteNvmePathRequest{
			Name: nvmeControllerToPathResourceID(id),
		})
		if err != nil {
			log.Printf("could not disconnect Remote Nvme path: %v", err)
			return err
		}

		response, err := client.DeleteNvmeRemoteController(ctx, &pb.DeleteNvmeRemoteControllerRequest{Name: id})
		if err != nil {
			log.Printf("could not disconnect Remote Nvme controller: %v", err)
			return err
		}
		log.Printf("disconnected: %v", response)
		return nil
	}
	log.Printf("Remote Nvme controller disconnected successfully")
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
				Name: subsystemID,
				Spec: &pb.NvmeSubsystemSpec{
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
			Parent: subsystemID,
			NvmeController: &pb.NvmeController{
				Name: controllerID,
				Spec: &pb.NvmeControllerSpec{
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

	client1 := pb.NewNullVolumeServiceClient(conn)
	response, err := client1.ListNullVolumes(ctx, &pb.ListNullVolumesRequest{})

	if err != nil {
		log.Println(err)
		return "", err
	}

	volumeData := response.NullVolumes
	volumeID := ""
	for _, data := range volumeData {
		uuid := strings.ReplaceAll(data.Uuid, "-", "")
		if uuid == nguid {
			volumeID = data.Name
		}
	}
	if volumeID == "" {
		return "", errors.New("volume ID not found")
	}

	client2 := pb.NewFrontendNvmeServiceClient(conn)
	resp, err := client2.CreateNvmeNamespace(ctx, &pb.CreateNvmeNamespaceRequest{
		Parent: subSystemID,
		NvmeNamespace: &pb.NvmeNamespace{
			Name: id,
			Spec: &pb.NvmeNamespaceSpec{
				VolumeNameRef: volumeID,
				HostNsid:      hostID,
			},
		},
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(resp)
	return resp.Name, nil
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

func nvmeControllerToPathResourceID(resourceID string) string {
	return resourceID + "path"
}
