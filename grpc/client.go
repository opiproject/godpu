// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

package grpc

import (
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type clientImpl struct {
	addr string // address of OPI gRPC server
	d    Dialler
}

// Dialler defines the function type that creates a gRPC connection
type Dialler func(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error)

// Closer defines the function type that closes gRPC connections
type Closer func()

type Connector interface {
	NewConn() (grpc.ClientConnInterface, Closer, error)
}

// New returns a new gRPC connector for the server at the given address
func New(address string) (Connector, error) {
	return NewWithDialler(address, grpc.Dial)
}

// NewWithDialler returns a new gRPC client for the server at the given address using the gRPC dialler provided
func NewWithDialler(address string, d Dialler) (Connector, error) {
	if len(address) == 0 {
		return nil, errors.New("cannot use empty address")
	}

	if d == nil {
		return nil, errors.New("grpc dialler is nil")
	}

	return clientImpl{
		addr: address,
		d:    d,
	}, nil
}

// NewConn creates a new gRPC connection
func (c clientImpl) NewConn() (grpc.ClientConnInterface, Closer, error) {
	conn, err := c.d(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	closer := func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("did not close connection: %v", err)
		}
	}
	return conn, closer, nil
}
