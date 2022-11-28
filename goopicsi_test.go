// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

package goopicsi

import (
	"context"
	"log"
	"testing"

	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"github.com/stretchr/testify/require"
	"go.nhat.io/grpcmock"
	"go.nhat.io/grpcmock/service"

	"github.com/stretchr/testify/assert"
)

func mockNVMfRemoteControllerServiceServer(m ...grpcmock.ServerOption) grpcmock.ServerMockerWithContextDialer {
	opts := []grpcmock.ServerOption{grpcmock.RegisterServiceFromMethods(service.Method{
		// Provide a service definition with request and response type.
		ServiceName: "opi_api.storage.v1.NVMfRemoteControllerService",
		MethodName:  "NVMfRemoteControllerGet",
		MethodType:  service.TypeUnary,
		Input:       &pb.NVMfRemoteControllerGetRequest{},
		Output:      &pb.NVMfRemoteControllerGetResponse{},
	})}
	opts = append(opts, m...)

	return grpcmock.MockServerWithBufConn(opts...)
}

func TestServer(t *testing.T) {
	t.Parallel()

	const getItem = "/opi_api.storage.v1.NVMfRemoteControllerService/NVMfRemoteControllerGet"

	testCases := []struct {
		scenario   string
		mockServer grpcmock.ServerMockerWithContextDialer
		request    pb.NVMfRemoteControllerGetRequest
		expected   pb.NVMfRemoteControllerGetResponse
	}{
		{
			scenario: "success",
			mockServer: mockNVMfRemoteControllerServiceServer(func(s *grpcmock.Server) {
				s.ExpectUnary(getItem).
					WithPayload(&pb.NVMfRemoteControllerGetRequest{Id: 12}).
					Return(&pb.NVMfRemoteControllerGetResponse{Ctrl: &pb.NVMfRemoteController{Subnqn: "nqn1"}})
			}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			_, dialer := tc.mockServer(t)

			// Use the dialer in your client, do the request and assertions.
			// For example:
			out := &pb.NVMfRemoteControllerGetResponse{}
			err := grpcmock.InvokeUnary(context.Background(),
				getItem, &pb.NVMfRemoteControllerGetRequest{Id: 12}, out,
				grpcmock.WithInsecure(),
				grpcmock.WithContextDialer(dialer),
			)

			require.NoError(t, err)

			assert.Equal(t, "nqn1", out.Ctrl.Subnqn)
			var c goopicsi

			res, err := c.goopicsiInterface.NVMeControllerGet(int64(12))
			println(res)

			// Server is closed at the end, and the ExpectationsWereMet() is also called, automatically!
		})
	}
}

func TestNVMeControllerConnect(t *testing.T) {
	c := goopicsi{}
	err := c.goopicsiInterface.NVMeControllerConnect(12, "", "", 44565)
	if err != nil {
		log.Println(err)
	}
	assert.Error(t, err)
}

func TestNVMeControllerList(t *testing.T) {
	c := goopicsi{}
	resp, err := c.goopicsiInterface.NVMeControllerList()
	if err != nil {
		log.Println(err)
	}
	log.Printf("NVMf Remote Connections: %v", resp)
}

func TestNVMeControllerGet(t *testing.T) {
	c := goopicsi{}
	resp, err := c.goopicsiInterface.NVMeControllerGet(12)
	if err != nil {
		log.Println(err)
	}
	log.Printf("NVMf remote connection corresponding to the ID: %v", resp)
}

func TestNVMeControllerDisconnect(t *testing.T) {
	c := goopicsi{}
	err := c.goopicsiInterface.NVMeControllerDisconnect(12)
	if err != nil {
		log.Println(err)
	}
}

func TestCreateNVMeNamespace(t *testing.T) {
	c := goopicsi{}
	resp, err := c.goopicsiInterface.CreateNVMeNamespace("1", "nqn", "nguid", 1)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

func TestDeleteNVMeNamespace(t *testing.T) {
	c := goopicsi{}
	err := c.goopicsiInterface.DeleteNVMeNamespace("1")
	if err != nil {
		log.Println(err)
	}
}

func TestExposeRemoteNVMe(t *testing.T) {
	c := goopicsi{}
	err := c.goopicsiInterface.ExposeRemoteNVMe("subsystem1", "nqn.2022-09.io.spdk:test", 10, "controller1")
	if err != nil {
		log.Println(err)
	}
}
