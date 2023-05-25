// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"fmt"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/opiproject/godpu/testing/mock-server/server"
	"github.com/opiproject/godpu/testing/mock-server/stub"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"

	"github.com/stretchr/testify/assert"
)

type GoopcsiTestSuite struct {
	suite.Suite
}

func (suite *GoopcsiTestSuite) SetupSuite() {
	RunServer()
}

// RunServer launches mock grpc server
func RunServer() {
	fmt.Println("RUNNING MOCK SERVER")
	const (
		csiAddress       = "localhost:50051"
		defaultStubsPath = "../testing/mock-server/stubs"
		apiPort          = "4771"
	)

	// run admin stub server
	stub.RunStubServer(stub.Options{
		StubPath: defaultStubsPath,
		Port:     apiPort,
		BindAddr: "0.0.0.0",
	})
	var protocol string
	if strings.Contains(csiAddress, ":") {
		protocol = "tcp"
	} else {
		protocol = "unix"
	}
	lis, err := net.Listen(protocol, csiAddress)
	if err != nil {
		fmt.Println(err, "failed to listen on address", "address", csiAddress)
		os.Exit(1)
	}

	MockServer := grpc.NewServer()

	pb.RegisterFrontendNvmeServiceServer(MockServer, &server.GoopCSI{})
	pb.RegisterNVMfRemoteControllerServiceServer(MockServer, &server.GoopCSI{})
	pb.RegisterNullDebugServiceServer(MockServer, &server.GoopCSI{})

	fmt.Printf("Serving gRPC on %s\n", csiAddress)
	errChan := make(chan error)

	// run blocking call in a separate goroutine, report errors via channel
	go func() {
		if err := MockServer.Serve(lis); err != nil {
			errChan <- err
		}
	}()
}

func (suite *GoopcsiTestSuite) TearDownTestSuite() {
	suite.T().Log("Cleaning up resources..")
}

func (suite *GoopcsiTestSuite) TestExposeRemoteNvme() {
	// Negative scenario
	subsystemID, controllerID, err := ExposeRemoteNvme("nqn.2022-09.io.spdk:test", 10)
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), subsystemID, "ExposeRemoteNvme failed")
	assert.Empty(suite.T(), controllerID, "ExposeRemoteNvme failed")
}

func (suite *GoopcsiTestSuite) TestCreateNvmeNamespace() {
	// scenario: when volume ID not found
	resp, err := CreateNvmeNamespace("1", "nqn", "nguid", 1)
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), resp, "CreateNvmeNamespace failed with invalid volume ID")
}

func (suite *GoopcsiTestSuite) TestNvmeControllerDisconnect() {
	// scenario: when connection already exists
	err := NvmeControllerDisconnect("12")
	assert.NoError(suite.T(), err)
}

func (suite *GoopcsiTestSuite) TestNvmeControllerConnect() {
	// scenario: when connection already exists
	err := NvmeControllerConnect("12", "", "", 44565, "")
	assert.NoError(suite.T(), err)
}

func (suite *GoopcsiTestSuite) TestNvmeControllerList() {
	resp, err := NvmeControllerList()
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), resp, "ListControllers success")
}

func (suite *GoopcsiTestSuite) TestNvmeControllerGet() {
	// positive scenario
	resp, err := NvmeControllerGet("12")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), resp, "GetController success")

	// negative scenario
	resp, err = NvmeControllerGet("invalid")
	assert.Error(suite.T(), err, "GetController failed")
	assert.Empty(suite.T(), resp, "GetController failed")
}

func (suite *GoopcsiTestSuite) TestDeleteNvmeNamespace() {
	// positive scenario
	err := DeleteNvmeNamespace("1")
	assert.NoError(suite.T(), err, "DeleteNvmeNamespace success")

	// negative scenario
	err = DeleteNvmeNamespace("invalid")
	assert.Error(suite.T(), err, "DeleteNvmeNamespace failed")
}

func (suite *GoopcsiTestSuite) TestGenerateHostNQN() {
	hostNQN := GenerateHostNQN()
	assert.NotNil(suite.T(), hostNQN, "GenerateHostNQN success")
}

func TestGoopcsiTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping as requested by short flag")
	}
	testSuite := new(GoopcsiTestSuite)
	suite.Run(t, testSuite)
	testSuite.TearDownTestSuite()
}
