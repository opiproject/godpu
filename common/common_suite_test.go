package common_test

import (
	"errors"
	"google.golang.org/grpc"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCommon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Common Suite")
}

func diallerNoError(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	dummyConn := grpc.ClientConn{}
	return &dummyConn, nil
}

func diallerWithError(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return nil, errors.New("error creating connection")

}