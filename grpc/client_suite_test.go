// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

package grpc_test

import (
	"errors"
	"google.golang.org/grpc"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCommon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gRPC Suite")
}

func diallerNoError(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	dummyConn := grpc.ClientConn{}
	return &dummyConn, nil
}

func diallerWithError(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return nil, errors.New("error creating connection")

}
