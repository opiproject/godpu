/*
Copyright © 2021-2022 Dell Inc. or its subsidiaries. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package server implements mock gRPC services
package server

import (
	"bytes"
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

// CreateNVMeSubsystem creates mock NVMe subsystem
func (s *GoopCSI) CreateNVMeSubsystem(ctx context.Context, request *pb.CreateNVMeSubsystemRequest) (*pb.NVMeSubsystem, error) {
	// TODO implement me
	panic("implement me")
}

// DeleteNVMeSubsystem deletes mock NVMe subsystem
func (s *GoopCSI) DeleteNVMeSubsystem(ctx context.Context, request *pb.DeleteNVMeSubsystemRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

// UpdateNVMeSubsystem updates a mock NVMe subsystem
func (s *GoopCSI) UpdateNVMeSubsystem(ctx context.Context, request *pb.UpdateNVMeSubsystemRequest) (*pb.NVMeSubsystem, error) {
	// TODO implement me
	panic("implement me")
}

// ListNVMeSubsystems lists mock NVMe subsystems
func (s *GoopCSI) ListNVMeSubsystems(ctx context.Context, request *pb.ListNVMeSubsystemsRequest) (*pb.ListNVMeSubsystemsResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetNVMeSubsystem gets a mock NVMe subsystem
func (s *GoopCSI) GetNVMeSubsystem(ctx context.Context, request *pb.GetNVMeSubsystemRequest) (*pb.NVMeSubsystem, error) {
	// TODO implement me
	panic("implement me")
}

// NVMeSubsystemStats gets mock subsystem stats
func (s *GoopCSI) NVMeSubsystemStats(ctx context.Context, request *pb.NVMeSubsystemStatsRequest) (*pb.NVMeSubsystemStatsResponse, error) {
	// TODO implement me
	panic("implement me")
}

// CreateNVMeController creates a mock NVMe controller
func (s *GoopCSI) CreateNVMeController(ctx context.Context, request *pb.CreateNVMeControllerRequest) (*pb.NVMeController, error) {
	// TODO implement me
	panic("implement me")
}

// DeleteNVMeController deletes a mock NVMe controller
func (s *GoopCSI) DeleteNVMeController(ctx context.Context, request *pb.DeleteNVMeControllerRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

// UpdateNVMeController updates a mock NVMe controller
func (s *GoopCSI) UpdateNVMeController(ctx context.Context, request *pb.UpdateNVMeControllerRequest) (*pb.NVMeController, error) {
	// TODO implement me
	panic("implement me")
}

// ListNVMeControllers lists mock controllers
func (s *GoopCSI) ListNVMeControllers(ctx context.Context, request *pb.ListNVMeControllersRequest) (*pb.ListNVMeControllersResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetNVMeController gets a mock NVMe controller
func (s *GoopCSI) GetNVMeController(ctx context.Context, request *pb.GetNVMeControllerRequest) (*pb.NVMeController, error) {
	// TODO implement me
	panic("implement me")
}

// NVMeControllerStats gets mock stats
func (s *GoopCSI) NVMeControllerStats(ctx context.Context, request *pb.NVMeControllerStatsRequest) (*pb.NVMeControllerStatsResponse, error) {
	// TODO implement me
	panic("implement me")
}

// CreateNVMeNamespace creates a mock NVMe namespace
func (s *GoopCSI) CreateNVMeNamespace(ctx context.Context, request *pb.CreateNVMeNamespaceRequest) (*pb.NVMeNamespace, error) {
	out := &pb.NVMeNamespace{}
	err := FindStub("FrontendNvmeServiceServer", "CreateNVMeNamespace", request, out)
	return out, err
}

// DeleteNVMeNamespace deletes a mock NVMe namespace
func (s *GoopCSI) DeleteNVMeNamespace(ctx context.Context, request *pb.DeleteNVMeNamespaceRequest) (*emptypb.Empty, error) {
	out := &emptypb.Empty{}
	err := FindStub("FrontendNvmeServiceServer", "DeleteNVMeNamespace", request, out)
	return out, err
}

// UpdateNVMeNamespace updates a mock namespace
func (s *GoopCSI) UpdateNVMeNamespace(ctx context.Context, request *pb.UpdateNVMeNamespaceRequest) (*pb.NVMeNamespace, error) {
	// TODO implement me
	panic("implement me")
}

// ListNVMeNamespaces lists mock namespaces
func (s *GoopCSI) ListNVMeNamespaces(ctx context.Context, request *pb.ListNVMeNamespacesRequest) (*pb.ListNVMeNamespacesResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetNVMeNamespace gets a mock namespace
func (s *GoopCSI) GetNVMeNamespace(ctx context.Context, request *pb.GetNVMeNamespaceRequest) (*pb.NVMeNamespace, error) {
	// TODO implement me
	panic("implement me")
}

// NVMeNamespaceStats gets mock namespace stats
func (s *GoopCSI) NVMeNamespaceStats(ctx context.Context, request *pb.NVMeNamespaceStatsRequest) (*pb.NVMeNamespaceStatsResponse, error) {
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
