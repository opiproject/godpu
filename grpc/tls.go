// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package grpc contains utility functions
package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc/credentials"
)

// TLSConfig contains information required to enable TLS for gRPC client.
type TLSConfig struct {
	ClientCertPath string
	ClientKeyPath  string
	CaCertPath     string
}

// ParseTLSFiles parses a string containing server certificate,
// server key and CA certificate separated by `:`
func ParseTLSFiles(tlsFiles string) (TLSConfig, error) {
	files := strings.Split(tlsFiles, ":")

	numOfFiles := len(files)
	if numOfFiles != 3 {
		return TLSConfig{}, errors.New("wrong number of path entries provided." +
			"Expect <server cert>:<server key>:<ca cert> are provided separated by `:`")
	}

	tls := TLSConfig{}

	const emptyPathErr = "empty %s path is not allowed"
	tls.ClientCertPath = files[0]
	if tls.ClientCertPath == "" {
		return TLSConfig{}, fmt.Errorf(emptyPathErr, "server cert")
	}

	tls.ClientKeyPath = files[1]
	if tls.ClientKeyPath == "" {
		return TLSConfig{}, fmt.Errorf(emptyPathErr, "server key")
	}

	tls.CaCertPath = files[2]
	if tls.CaCertPath == "" {
		return TLSConfig{}, fmt.Errorf(emptyPathErr, "CA cert")
	}

	return tls, nil
}

// SetupTLSCredentials returns a service options to enable TLS for gRPC client
func SetupTLSCredentials(config TLSConfig) (credentials.TransportCredentials, error) {
	return setupTLSCredentials(config, tls.LoadX509KeyPair, os.ReadFile)
}

func setupTLSCredentials(config TLSConfig,
	loadX509KeyPair func(string, string) (tls.Certificate, error),
	readFile func(string) ([]byte, error),
) (credentials.TransportCredentials, error) {
	clientCert, err := loadX509KeyPair(config.ClientCertPath, config.ClientKeyPath)
	if err != nil {
		return nil, err
	}
	c := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		MinVersion:   tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}

	c.RootCAs = x509.NewCertPool()
	log.Println("Loading client ca certificate:", config.CaCertPath)

	clientCaCert, err := readFile(config.CaCertPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate: %v. error: %v", config.CaCertPath, err)
	}

	if !c.RootCAs.AppendCertsFromPEM(clientCaCert) {
		return nil, fmt.Errorf("failed to add client CA's certificate: %v", config.CaCertPath)
	}

	return credentials.NewTLS(c), nil
}
