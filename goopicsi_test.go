// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

package goopicsi

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/opiproject/goopicsi/test/mock-server/server"
	"github.com/opiproject/goopicsi/test/mock-server/stub"
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
		defaultStubsPath = "test/mock-server/stubs"
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

// TODO These test cases should be reverted with mock server implementation
// func TestNVMeControllerConnect(t *testing.T) {
// 	err := NVMeControllerConnect("12", "", "", 44565, "")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	assert.Error(t, err)
// }

// func TestNVMeControllerList(t *testing.T) {
// 	resp, err := NVMeControllerList()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Printf("NVMf Remote Connections: %v", resp)
// }

// func TestNVMeControllerGet(t *testing.T) {
// 	resp, err := NVMeControllerGet("12")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Printf("NVMf remote connection corresponding to the ID: %v", resp)
// }

// func TestNVMeControllerDisconnect(t *testing.T) {
// 	err := NVMeControllerDisconnect("12")
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func TestCreateNVMeNamespace(t *testing.T) {
// 	resp, err := CreateNVMeNamespace("1", "nqn", "nguid", 1)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(resp)
// }

// func TestExposeRemoteNVMe(t *testing.T) {
// 	subsystemID, controllerID, err := ExposeRemoteNVMe("nqn.2022-09.io.spdk:test", 10)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Printf("Subsystem ID: %s", subsystemID)
// 	log.Printf("Controller Id: %s", controllerID)
// }

func (suite *GoopcsiTestSuite) TestDeleteNVMeNamespace() {
	// positive scenario
	err := DeleteNVMeNamespace("1")
	assert.NoError(suite.T(), err, "DeleteNVMeNamespace success")

	// negative scenario
	err = DeleteNVMeNamespace("invalid")
	assert.Error(suite.T(), err, "DeleteNVMeNamespace failed")
}

func TestGenerateHostNQN(t *testing.T) {
	hostNQN := GenerateHostNQN()
	log.Println(hostNQN)
}

func TestGoopcsiTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping as requested by short flag")
	}
	testSuite := new(GoopcsiTestSuite)
	suite.Run(t, testSuite)
	testSuite.TearDownTestSuite()
}
