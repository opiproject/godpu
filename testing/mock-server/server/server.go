// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2021-2022 Dell Inc, or its subsidiaries.

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
type GoopCSI struct{}

// CreateNullDebug creates a mock NullDebug
func (s *GoopCSI) CreateNullDebug(_ context2.Context, _ *pb.CreateNullDebugRequest) (*pb.NullDebug, error) {
	// TODO implement me
	panic("implement me")
}

// DeleteNullDebug Deletes a mock NullDebug
func (s *GoopCSI) DeleteNullDebug(_ context2.Context, _ *pb.DeleteNullDebugRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

// UpdateNullDebug Updated mock NullDebug
func (s *GoopCSI) UpdateNullDebug(_ context2.Context, _ *pb.UpdateNullDebugRequest) (*pb.NullDebug, error) {
	// TODO implement me
	panic("implement me")
}

// ListNullDebugs Lists mock nullDebugs
func (s *GoopCSI) ListNullDebugs(_ context2.Context, request *pb.ListNullDebugsRequest) (*pb.ListNullDebugsResponse, error) {
	out := &pb.ListNullDebugsResponse{}
	err := FindStub("NullDebugServiceServer", "ListNullDebugs", request, out)
	return out, err
}

// GetNullDebug Gets mock NullDebug
func (s *GoopCSI) GetNullDebug(_ context2.Context, _ *pb.GetNullDebugRequest) (*pb.NullDebug, error) {
	// TODO implement me
	panic("implement me")
}

// NullDebugStats gets mock NullDebugStats
func (s *GoopCSI) NullDebugStats(_ context2.Context, _ *pb.NullDebugStatsRequest) (*pb.NullDebugStatsResponse, error) {
	// TODO implement me
	panic("implement me")
}

// CreateNVMfRemoteController creates a mock NVMf Remote controller
func (s *GoopCSI) CreateNVMfRemoteController(_ context2.Context, request *pb.CreateNVMfRemoteControllerRequest) (*pb.NVMfRemoteController, error) {
	out := &pb.NVMfRemoteController{}
	err := FindStub("NVMfRemoteControllerServiceServer", "CreateNVMfRemoteController", request, out)
	return out, err
}

// DeleteNVMfRemoteController deletes a mock NVMfRemote Controller
func (s *GoopCSI) DeleteNVMfRemoteController(_ context2.Context, request *pb.DeleteNVMfRemoteControllerRequest) (*emptypb.Empty, error) {
	out := &emptypb.Empty{}
	err := FindStub("NVMfRemoteControllerServiceServer", "DeleteNVMfRemoteController", request, out)
	return out, err
}

// UpdateNVMfRemoteController updates mock NVMf Remote Controller
func (s *GoopCSI) UpdateNVMfRemoteController(_ context2.Context, _ *pb.UpdateNVMfRemoteControllerRequest) (*pb.NVMfRemoteController, error) {
	// TODO implement me
	panic("implement me")
}

// ListNVMfRemoteControllers Lists mock NVMfRemote Controllers
func (s *GoopCSI) ListNVMfRemoteControllers(_ context2.Context, request *pb.ListNVMfRemoteControllersRequest) (*pb.ListNVMfRemoteControllersResponse, error) {
	out := &pb.ListNVMfRemoteControllersResponse{}
	err := FindStub("NVMfRemoteControllerServiceServer", "ListNVMfRemoteControllers", request, out)
	return out, err
}

// GetNVMfRemoteController Gets an NVMf Remote controller
func (s *GoopCSI) GetNVMfRemoteController(_ context2.Context, request *pb.GetNVMfRemoteControllerRequest) (*pb.NVMfRemoteController, error) {
	out := &pb.NVMfRemoteController{}
	err := FindStub("NVMfRemoteControllerServiceServer", "GetNVMfRemoteController", request, out)
	return out, err
}

// NVMfRemoteControllerReset Resets mock Remote Controller
func (s *GoopCSI) NVMfRemoteControllerReset(_ context2.Context, _ *pb.NVMfRemoteControllerResetRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

// NVMfRemoteControllerStats gets mock stats
func (s *GoopCSI) NVMfRemoteControllerStats(_ context2.Context, _ *pb.NVMfRemoteControllerStatsRequest) (*pb.NVMfRemoteControllerStatsResponse, error) {
	// TODO implement me
	panic("implement me")
}

// CreateNVMeSubsystem creates mock NVMe subsystem
func (s *GoopCSI) CreateNVMeSubsystem(_ context.Context, request *pb.CreateNVMeSubsystemRequest) (*pb.NVMeSubsystem, error) {
	out := &pb.NVMeSubsystem{}
	err := FindStub("FrontendNvmeServiceServer", "CreateNVMeSubsystem", request, out)
	return out, err
}

// DeleteNVMeSubsystem deletes mock NVMe subsystem
func (s *GoopCSI) DeleteNVMeSubsystem(_ context.Context, _ *pb.DeleteNVMeSubsystemRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

// UpdateNVMeSubsystem updates a mock NVMe subsystem
func (s *GoopCSI) UpdateNVMeSubsystem(_ context.Context, _ *pb.UpdateNVMeSubsystemRequest) (*pb.NVMeSubsystem, error) {
	// TODO implement me
	panic("implement me")
}

// ListNVMeSubsystems lists mock NVMe subsystems
func (s *GoopCSI) ListNVMeSubsystems(_ context.Context, _ *pb.ListNVMeSubsystemsRequest) (*pb.ListNVMeSubsystemsResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetNVMeSubsystem gets a mock NVMe subsystem
func (s *GoopCSI) GetNVMeSubsystem(_ context.Context, request *pb.GetNVMeSubsystemRequest) (*pb.NVMeSubsystem, error) {
	out := &pb.NVMeSubsystem{}
	err := FindStub("FrontendNvmeServiceServer", "GetNVMeSubsystem", request, out)
	return out, err
}

// NVMeSubsystemStats gets mock subsystem stats
func (s *GoopCSI) NVMeSubsystemStats(_ context.Context, _ *pb.NVMeSubsystemStatsRequest) (*pb.NVMeSubsystemStatsResponse, error) {
	// TODO implement me
	panic("implement me")
}

// CreateNVMeController creates a mock NVMe controller
func (s *GoopCSI) CreateNVMeController(_ context.Context, _ *pb.CreateNVMeControllerRequest) (*pb.NVMeController, error) {
	// TODO implement me
	panic("implement me")
}

// DeleteNVMeController deletes a mock NVMe controller
func (s *GoopCSI) DeleteNVMeController(_ context.Context, _ *pb.DeleteNVMeControllerRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

// UpdateNVMeController updates a mock NVMe controller
func (s *GoopCSI) UpdateNVMeController(_ context.Context, _ *pb.UpdateNVMeControllerRequest) (*pb.NVMeController, error) {
	// TODO implement me
	panic("implement me")
}

// ListNVMeControllers lists mock controllers
func (s *GoopCSI) ListNVMeControllers(_ context.Context, _ *pb.ListNVMeControllersRequest) (*pb.ListNVMeControllersResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetNVMeController gets a mock NVMe controller
func (s *GoopCSI) GetNVMeController(_ context.Context, request *pb.GetNVMeControllerRequest) (*pb.NVMeController, error) {
	out := &pb.NVMeController{}
	err := FindStub("FrontendNvmeServiceServer", "GetNVMeController", request, out)
	return out, err
}

// NVMeControllerStats gets mock stats
func (s *GoopCSI) NVMeControllerStats(_ context.Context, _ *pb.NVMeControllerStatsRequest) (*pb.NVMeControllerStatsResponse, error) {
	// TODO implement me
	panic("implement me")
}

// CreateNVMeNamespace creates a mock NVMe namespace
func (s *GoopCSI) CreateNVMeNamespace(_ context.Context, request *pb.CreateNVMeNamespaceRequest) (*pb.NVMeNamespace, error) {
	out := &pb.NVMeNamespace{}
	err := FindStub("FrontendNvmeServiceServer", "CreateNVMeNamespace", request, out)
	return out, err
}

// DeleteNVMeNamespace deletes a mock NVMe namespace
func (s *GoopCSI) DeleteNVMeNamespace(_ context.Context, request *pb.DeleteNVMeNamespaceRequest) (*emptypb.Empty, error) {
	out := &emptypb.Empty{}
	err := FindStub("FrontendNvmeServiceServer", "DeleteNVMeNamespace", request, out)
	return out, err
}

// UpdateNVMeNamespace updates a mock namespace
func (s *GoopCSI) UpdateNVMeNamespace(_ context.Context, _ *pb.UpdateNVMeNamespaceRequest) (*pb.NVMeNamespace, error) {
	// TODO implement me
	panic("implement me")
}

// ListNVMeNamespaces lists mock namespaces
func (s *GoopCSI) ListNVMeNamespaces(_ context.Context, _ *pb.ListNVMeNamespacesRequest) (*pb.ListNVMeNamespacesResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetNVMeNamespace gets a mock namespace
func (s *GoopCSI) GetNVMeNamespace(_ context.Context, _ *pb.GetNVMeNamespaceRequest) (*pb.NVMeNamespace, error) {
	// TODO implement me
	panic("implement me")
}

// NVMeNamespaceStats gets mock namespace stats
func (s *GoopCSI) NVMeNamespaceStats(_ context.Context, _ *pb.NVMeNamespaceStatsRequest) (*pb.NVMeNamespaceStatsResponse, error) {
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
