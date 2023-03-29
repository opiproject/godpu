/* SPDX-License-Identifier: Apache-2.0
   Copyright (c) 2023 Dell Inc, or its subsidiaries.
*/

package common

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

type Client interface {
	NewGrpcConn() (grpc.ClientConnInterface, Closer, error)
}

// NewClient returns a new gRPC client for the server at the given address
func NewClient(address string) (Client, error) {
	return NewClientWithDialler(address, grpc.Dial)
}

// NewClientWithDialler returns a new gRPC client for the server at the given address using the gRPC dialler provided
func NewClientWithDialler(address string, d Dialler) (Client, error) {
	if len(address) == 0 {
		return nil, errors.New("cannot use empty address")
	}
	return clientImpl{
		addr: address,
		d:    d,
	}, nil
}

// NewGrpcConn creates a new gRPC connection
func (c clientImpl) NewGrpcConn() (grpc.ClientConnInterface, Closer, error) {
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
