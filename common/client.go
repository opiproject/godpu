/* SPDX-License-Identifier: Apache-2.0
   Copyright (c) 2023 Dell Inc, or its subsidiaries.
*/

package common

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type clientImpl struct {
	addr string // address of OPI gRPC server
	d    dialler
}

type dialler func(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error)

type Closer func()

type GetGrpcConn func(addr string) (grpc.ClientConnInterface, Closer, error)

type Client interface {
	NewGrpcConn() (grpc.ClientConnInterface, Closer, error)
}

func NewClient(addr string) Client {
	return clientImpl{
		addr: addr,
		d:    grpc.Dial,
	}
}

func (c clientImpl) NewGrpcConn() (grpc.ClientConnInterface, Closer, error) {
	conn, err := grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
