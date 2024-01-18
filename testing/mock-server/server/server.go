// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2021-2023 Dell Inc, or its subsidiaries.
// Copyright (C) 2023 Intel Corporation

// Package server implements mock gRPC services
package server

import (
	"bytes"
	context2 "context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
)

// GoopCSI mock gRPC server to implement mock service calls
type GoopCSI struct {
	pb.UnimplementedNullVolumeServiceServer
	pb.UnimplementedNvmeRemoteControllerServiceServer
	pb.UnimplementedFrontendNvmeServiceServer
}

var _ pb.NullVolumeServiceServer = &GoopCSI{}
var _ pb.NvmeRemoteControllerServiceServer = &GoopCSI{}
var _ pb.FrontendNvmeServiceServer = &GoopCSI{}

// CreateNullVolume creates a mock NullVolume
func (s *GoopCSI) CreateNullVolume(_ context2.Context, _ *pb.CreateNullVolumeRequest) (*pb.NullVolume, error) {
	// TODO implement me
	panic("implement me")
}

// DeleteNullVolume Deletes a mock NullVolume
func (s *GoopCSI) DeleteNullVolume(_ context2.Context, _ *pb.DeleteNullVolumeRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

// UpdateNullVolume Updated mock NullVolume
func (s *GoopCSI) UpdateNullVolume(_ context2.Context, _ *pb.UpdateNullVolumeRequest) (*pb.NullVolume, error) {
	// TODO implement me
	panic("implement me")
}

// ListNullVolumes Lists mock nullDebugs
func (s *GoopCSI) ListNullVolumes(_ context2.Context, request *pb.ListNullVolumesRequest) (*pb.ListNullVolumesResponse, error) {
	out := &pb.ListNullVolumesResponse{}
	err := FindStub("NullVolumeServiceServer", "ListNullVolumes", request, out)
	return out, err
}

// GetNullVolume Gets mock NullVolume
func (s *GoopCSI) GetNullVolume(_ context2.Context, _ *pb.GetNullVolumeRequest) (*pb.NullVolume, error) {
	// TODO implement me
	panic("implement me")
}

// StatsNullVolume gets mock StatsNullVolume
func (s *GoopCSI) StatsNullVolume(_ context2.Context, _ *pb.StatsNullVolumeRequest) (*pb.StatsNullVolumeResponse, error) {
	// TODO implement me
	panic("implement me")
}

// CreateNvmeRemoteController creates a mock Nvme Remote controller
func (s *GoopCSI) CreateNvmeRemoteController(_ context2.Context, request *pb.CreateNvmeRemoteControllerRequest) (*pb.NvmeRemoteController, error) {
	out := &pb.NvmeRemoteController{}
	err := FindStub("NvmeRemoteControllerServiceServer", "CreateNvmeRemoteController", request, out)
	return out, err
}

// DeleteNvmeRemoteController deletes a mock NvmeRemote Controller
func (s *GoopCSI) DeleteNvmeRemoteController(_ context2.Context, request *pb.DeleteNvmeRemoteControllerRequest) (*emptypb.Empty, error) {
	out := &emptypb.Empty{}
	err := FindStub("NvmeRemoteControllerServiceServer", "DeleteNvmeRemoteController", request, out)
	return out, err
}

// UpdateNvmeRemoteController updates mock Nvme Remote Controller
func (s *GoopCSI) UpdateNvmeRemoteController(_ context2.Context, _ *pb.UpdateNvmeRemoteControllerRequest) (*pb.NvmeRemoteController, error) {
	// TODO implement me
	panic("implement me")
}

// ListNvmeRemoteControllers Lists mock NvmeRemote Controllers
func (s *GoopCSI) ListNvmeRemoteControllers(_ context2.Context, request *pb.ListNvmeRemoteControllersRequest) (*pb.ListNvmeRemoteControllersResponse, error) {
	out := &pb.ListNvmeRemoteControllersResponse{}
	err := FindStub("NvmeRemoteControllerServiceServer", "ListNvmeRemoteControllers", request, out)
	return out, err
}

// GetNvmeRemoteController Gets an Nvme Remote controller
func (s *GoopCSI) GetNvmeRemoteController(_ context2.Context, request *pb.GetNvmeRemoteControllerRequest) (*pb.NvmeRemoteController, error) {
	out := &pb.NvmeRemoteController{}
	err := FindStub("NvmeRemoteControllerServiceServer", "GetNvmeRemoteController", request, out)
	return out, err
}

// ResetNvmeRemoteController Resets mock Remote Controller
func (s *GoopCSI) ResetNvmeRemoteController(_ context2.Context, _ *pb.ResetNvmeRemoteControllerRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

// StatsNvmeRemoteController gets mock stats
func (s *GoopCSI) StatsNvmeRemoteController(_ context2.Context, _ *pb.StatsNvmeRemoteControllerRequest) (*pb.StatsNvmeRemoteControllerResponse, error) {
	// TODO implement me
	panic("implement me")
}

// ListNvmeRemoteNamespaces lists mock remote nvme namespaces
func (s *GoopCSI) ListNvmeRemoteNamespaces(context2.Context, *pb.ListNvmeRemoteNamespacesRequest) (*pb.ListNvmeRemoteNamespacesResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetNvmeRemoteNamespace gets mock remote nvme namespace
func (s *GoopCSI) GetNvmeRemoteNamespace(_ context2.Context, _ *pb.GetNvmeRemoteNamespaceRequest) (*pb.NvmeRemoteNamespace, error) {
	// TODO implement me
	panic("implement me")
}

// CreateNvmePath creates mock nvme path
func (s *GoopCSI) CreateNvmePath(_ context2.Context, request *pb.CreateNvmePathRequest) (*pb.NvmePath, error) {
	out := &pb.NvmePath{}
	err := FindStub("NvmeRemoteControllerServiceServer", "CreateNvmePath", request, out)
	return out, err
}

// DeleteNvmePath deletes mock nvme path
func (s *GoopCSI) DeleteNvmePath(_ context2.Context, request *pb.DeleteNvmePathRequest) (*emptypb.Empty, error) {
	out := &emptypb.Empty{}
	err := FindStub("NvmeRemoteControllerServiceServer", "DeleteNvmePath", request, out)
	return out, err
}

// UpdateNvmePath updates mock Nvme Remote Path
func (s *GoopCSI) UpdateNvmePath(_ context2.Context, _ *pb.UpdateNvmePathRequest) (*pb.NvmePath, error) {
	// TODO implement me
	panic("implement me")
}

// ListNvmePaths Lists mock NvmeRemote Paths
func (s *GoopCSI) ListNvmePaths(_ context2.Context, _ *pb.ListNvmePathsRequest) (*pb.ListNvmePathsResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetNvmePath Gets an Nvme Remote Path
func (s *GoopCSI) GetNvmePath(_ context2.Context, _ *pb.GetNvmePathRequest) (*pb.NvmePath, error) {
	// TODO implement me
	panic("implement me")
}

// StatsNvmePath gets mock stats
func (s *GoopCSI) StatsNvmePath(_ context2.Context, _ *pb.StatsNvmePathRequest) (*pb.StatsNvmePathResponse, error) {
	// TODO implement me
	panic("implement me")
}

// CreateNvmeSubsystem creates mock Nvme subsystem
func (s *GoopCSI) CreateNvmeSubsystem(_ context.Context, request *pb.CreateNvmeSubsystemRequest) (*pb.NvmeSubsystem, error) {
	out := &pb.NvmeSubsystem{}
	err := FindStub("FrontendNvmeServiceServer", "CreateNvmeSubsystem", request, out)
	return out, err
}

// DeleteNvmeSubsystem deletes mock Nvme subsystem
func (s *GoopCSI) DeleteNvmeSubsystem(_ context.Context, _ *pb.DeleteNvmeSubsystemRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

// UpdateNvmeSubsystem updates a mock Nvme subsystem
func (s *GoopCSI) UpdateNvmeSubsystem(_ context.Context, _ *pb.UpdateNvmeSubsystemRequest) (*pb.NvmeSubsystem, error) {
	// TODO implement me
	panic("implement me")
}

// ListNvmeSubsystems lists mock Nvme subsystems
func (s *GoopCSI) ListNvmeSubsystems(_ context.Context, _ *pb.ListNvmeSubsystemsRequest) (*pb.ListNvmeSubsystemsResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetNvmeSubsystem gets a mock Nvme subsystem
func (s *GoopCSI) GetNvmeSubsystem(_ context.Context, request *pb.GetNvmeSubsystemRequest) (*pb.NvmeSubsystem, error) {
	out := &pb.NvmeSubsystem{}
	err := FindStub("FrontendNvmeServiceServer", "GetNvmeSubsystem", request, out)
	return out, err
}

// StatsNvmeSubsystem gets mock subsystem stats
func (s *GoopCSI) StatsNvmeSubsystem(_ context.Context, _ *pb.StatsNvmeSubsystemRequest) (*pb.StatsNvmeSubsystemResponse, error) {
	// TODO implement me
	panic("implement me")
}

// CreateNvmeController creates a mock Nvme controller
func (s *GoopCSI) CreateNvmeController(_ context.Context, _ *pb.CreateNvmeControllerRequest) (*pb.NvmeController, error) {
	// TODO implement me
	panic("implement me")
}

// DeleteNvmeController deletes a mock Nvme controller
func (s *GoopCSI) DeleteNvmeController(_ context.Context, _ *pb.DeleteNvmeControllerRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

// UpdateNvmeController updates a mock Nvme controller
func (s *GoopCSI) UpdateNvmeController(_ context.Context, _ *pb.UpdateNvmeControllerRequest) (*pb.NvmeController, error) {
	// TODO implement me
	panic("implement me")
}

// ListNvmeControllers lists mock controllers
func (s *GoopCSI) ListNvmeControllers(_ context.Context, _ *pb.ListNvmeControllersRequest) (*pb.ListNvmeControllersResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetNvmeController gets a mock Nvme controller
func (s *GoopCSI) GetNvmeController(_ context.Context, request *pb.GetNvmeControllerRequest) (*pb.NvmeController, error) {
	out := &pb.NvmeController{}
	err := FindStub("FrontendNvmeServiceServer", "GetNvmeController", request, out)
	return out, err
}

// StatsNvmeController gets mock stats
func (s *GoopCSI) StatsNvmeController(_ context.Context, _ *pb.StatsNvmeControllerRequest) (*pb.StatsNvmeControllerResponse, error) {
	// TODO implement me
	panic("implement me")
}

// CreateNvmeNamespace creates a mock Nvme namespace
func (s *GoopCSI) CreateNvmeNamespace(_ context.Context, request *pb.CreateNvmeNamespaceRequest) (*pb.NvmeNamespace, error) {
	out := &pb.NvmeNamespace{}
	err := FindStub("FrontendNvmeServiceServer", "CreateNvmeNamespace", request, out)
	return out, err
}

// DeleteNvmeNamespace deletes a mock Nvme namespace
func (s *GoopCSI) DeleteNvmeNamespace(_ context.Context, request *pb.DeleteNvmeNamespaceRequest) (*emptypb.Empty, error) {
	out := &emptypb.Empty{}
	err := FindStub("FrontendNvmeServiceServer", "DeleteNvmeNamespace", request, out)
	return out, err
}

// UpdateNvmeNamespace updates a mock namespace
func (s *GoopCSI) UpdateNvmeNamespace(_ context.Context, _ *pb.UpdateNvmeNamespaceRequest) (*pb.NvmeNamespace, error) {
	// TODO implement me
	panic("implement me")
}

// ListNvmeNamespaces lists mock namespaces
func (s *GoopCSI) ListNvmeNamespaces(_ context.Context, _ *pb.ListNvmeNamespacesRequest) (*pb.ListNvmeNamespacesResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetNvmeNamespace gets a mock namespace
func (s *GoopCSI) GetNvmeNamespace(_ context.Context, _ *pb.GetNvmeNamespaceRequest) (*pb.NvmeNamespace, error) {
	// TODO implement me
	panic("implement me")
}

// StatsNvmeNamespace gets mock namespace stats
func (s *GoopCSI) StatsNvmeNamespace(_ context.Context, _ *pb.StatsNvmeNamespaceRequest) (*pb.StatsNvmeNamespaceResponse, error) {
	// TODO implement me
	panic("implement me")
}

type payload struct {
	Service string      `json:"service"`
	Method  string      `json:"method"`
	Data    interface{} `json:"data"`
}

type response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// FindStub makes request to mock grpc server
func FindStub(service, method string, in, out interface{}) error {
	url := "http://localhost:4771/find"
	pyl := payload{
		Service: service,
		Method:  method,
		Data:    in,
	}
	byt, err := json.Marshal(pyl)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(byt)
	resp, err := http.DefaultClient.Post(url, "application/json", reader)
	if err != nil {
		return fmt.Errorf("error request to stub server %v", err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf(string(body))
	}

	respRPC := new(response)
	err = json.NewDecoder(resp.Body).Decode(respRPC)
	if err != nil {
		return fmt.Errorf("decoding json response %v", err)
	}

	if respRPC.Error != "" {
		return fmt.Errorf(respRPC.Error)
	}

	data, _ := json.Marshal(respRPC.Data)
	return json.Unmarshal(data, out)
}
