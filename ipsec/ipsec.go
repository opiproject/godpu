// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package ipsec implements the go library for OPI to be used to establish ipsec
package ipsec

import (
	"context"
	"log"
	"time"

	pb "github.com/opiproject/opi-api/security/v1/gen/go"
	probing "github.com/prometheus-community/pro-bing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn *grpc.ClientConn
)

// Stats returns statistics information from DPUs regaridng IPSEC
func Stats(address string) error {
	if conn == nil {
		err := dialConnection(address)
		if err != nil {
			return err
		}
	}

	client := pb.NewIPsecServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getStats(ctx, client)

	defer disconnectConnection()
	return nil
}

// TestIpsec runs few basic tests establishing ipsec tunnels, version and stats
func TestIpsec(address string, pingaddr string) error {
	// connection
	if conn == nil {
		err := dialConnection(address)
		if err != nil {
			return err
		}
	}
	// context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// IPsec
	c1 := pb.NewIPsecServiceClient(conn)

	// Print info
	getVersion(ctx, c1)
	getStats(ctx, c1)

	// Load IPsec connection
	loadConnections(ctx, c1)

	// Bring the connection up
	initConn := pb.IPsecInitiateRequest{
		Ike:   "opi-test",
		Child: "opi-child",
	}
	initRet, err := c1.IPsecInitiate(ctx, &initConn)
	if err != nil {
		log.Panicf("could not initiate IPsec tunnel: %v", err)
	}
	log.Printf("Initiated: %v", initRet)

	// List the ikeSas
	ikeSas := pb.IPsecListSasRequest{
		Ike: "opi-test",
	}
	listSasRet, err := c1.IPsecListSas(ctx, &ikeSas)
	if err != nil {
		log.Panicf("could not list ikeSas: %v", err)
	}
	log.Printf("Returned ikeSas: %v", listSasRet)

	// print various information
	listConnections(ctx, c1)
	listCertificates(ctx, c1)

	// Ping across the tunnel.
	doPing(pingaddr)

	// Rekey the IKE_SA
	rekeyConn := pb.IPsecRekeyRequest{
		Ike: "opi-test",
	}
	rekeyRet, err := c1.IPsecRekey(ctx, &rekeyConn)
	if err != nil {
		log.Panicf("could not rekey IPsec tunnel: %v", err)
	}
	log.Printf("Rekeyed IKE_SA %s: %v", "opi-test", rekeyRet)

	doCleanup(ctx, c1)
	defer disconnectConnection()
	return nil
}

func doCleanup(ctx context.Context, client pb.IPsecServiceClient) {
	// Terminate the connection
	termConn := pb.IPsecTerminateRequest{
		Ike: "opi-test",
	}

	termRet, err := client.IPsecTerminate(ctx, &termConn)
	if err != nil {
		log.Fatalf("could not terminate IPsec tunnel: %v", err)
	}
	log.Printf("Terminate: %v", termRet)

	// Unload
	unloadIpsec := pb.IPsecUnloadConnRequest{
		Name: "opi-test",
	}

	rs2, err := client.IPsecUnloadConn(ctx, &unloadIpsec)
	if err != nil {
		log.Fatalf("could not unload IPsec tunnel: %v", err)
	}
	log.Printf("Unloaded: %v", rs2)
}

func listConnections(ctx context.Context, client pb.IPsecServiceClient) {
	// List the connections
	listConn := pb.IPsecListConnsRequest{
		Ike: "opi-test",
	}
	listConnsRet, err := client.IPsecListConns(ctx, &listConn)
	if err != nil {
		log.Fatalf("could not list connections: %v", err)
	}
	log.Printf("Returned connections: %v", listConnsRet)
}

func listCertificates(ctx context.Context, client pb.IPsecServiceClient) {
	// List the certificates
	listCerts := pb.IPsecListCertsRequest{
		Type: "any",
	}
	listCertsRet, err := client.IPsecListCerts(ctx, &listCerts)
	if err != nil {
		log.Fatalf("could not list certificates: %v", err)
	}
	log.Printf("Returned certificates: %v", listCertsRet)
}

func getStats(ctx context.Context, client pb.IPsecServiceClient) {
	statsResp, err := client.IPsecStats(ctx, &pb.IPsecStatsRequest{})
	if err != nil {
		log.Fatalf("could not get IPsec stats")
	}
	log.Printf("IPsec stats\n%s", statsResp.GetStatus())
}

func getVersion(ctx context.Context, client pb.IPsecServiceClient) {
	vresp, err := client.IPsecVersion(ctx, &pb.IPsecVersionRequest{})
	if err != nil {
		log.Fatalf("could not get IPsec version")
	}
	log.Printf("Daemon  [%v]", vresp.GetDaemon())
	log.Printf("Version [%v]", vresp.GetVersion())
	log.Printf("Sysname [%v]", vresp.GetSysname())
	log.Printf("Release [%v]", vresp.GetRelease())
	log.Printf("Machine [%v]", vresp.GetMachine())
}

func doPing(a string) {
	// .NOTE: The container this test runs in is linked to the appropriate
	//        strongSwan container.
	pinger, err := probing.NewPinger(a)
	if err != nil {
		log.Fatalf("Cannot create Pinger")
	}
	pinger.Count = 5
	// .NOTE: This blocks until it finishes
	err = pinger.Run()
	if err != nil {
		log.Fatalf("Ping command to host 10.3.0.1 failed")
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats

	log.Printf("Ping stats: %v", stats)
}

func loadConnections(ctx context.Context, client pb.IPsecServiceClient) {
	localIpsec := pb.IPsecLoadConnRequest{
		Connection: &pb.Connection{
			Name:    "opi-test",
			Version: "2",
			Vips:    &pb.Vips{Vip: []string{"0.0.0.0"}},
			LocalAddrs: []*pb.Addrs{
				{
					Addr: "192.168.200.200",
				},
			},
			RemoteAddrs: []*pb.Addrs{
				{
					Addr: "192.168.200.210",
				},
			},
			LocalAuth:  &pb.LocalAuth{Auth: pb.AuthType_AUTH_TYPE_PSK, Id: "hacker@strongswan.org"},
			RemoteAuth: &pb.RemoteAuth{Auth: pb.AuthType_AUTH_TYPE_PSK, Id: "server.strongswan.org"},
			Children: []*pb.Child{
				{
					Name: "opi-child",
					EspProposals: &pb.Proposals{
						CryptoAlg: []pb.CryptoAlgorithm{pb.CryptoAlgorithm_CRYPTO_ALGORITHM_AES256GCM128},
						IntegAlg:  []pb.IntegAlgorithm{pb.IntegAlgorithm_INTEG_ALGORITHM_SHA512},
						Dhgroups:  []pb.DHGroups{pb.DHGroups_DH_GROUPS_CURVE25519},
					},
					RemoteTs: &pb.TrafficSelectors{
						Ts: []*pb.TrafficSelectors_TrafficSelector{
							{
								Cidr: "10.1.0.0/16",
							},
						},
					},
				},
			},
		},
	}
	rs1, err := client.IPsecLoadConn(ctx, &localIpsec)
	if err != nil {
		log.Fatalf("could not load IPsec tunnel: %v", err)
	}
	log.Printf("Loaded: %v", rs1)
}

func dialConnection(address string) error {
	var err error
	conn, err = grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect: %v", err)
		return err
	}
	return nil
}

func disconnectConnection() {
	err := conn.Close()
	if err != nil {
		log.Fatalf("Failed to close connection: %v", err)
	}
	log.Println("GRPC connection closed successfully")
}
