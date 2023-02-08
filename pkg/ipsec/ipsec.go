// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package ipsec implements the go library for OPI to be used to establish ipsec
package ipsec

import (
	"context"
	"log"
	"time"

	"github.com/go-ping/ping"
	pb "github.com/opiproject/opi-api/security/v1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn     *grpc.ClientConn
	address  = "localhost:50051"
	pingaddr = "localhost"
)

// Stats returns statistics information from DPUs regaridng IPSEC
func Stats() error {
	if conn == nil {
		err := dialConnection()
		if err != nil {
			return err
		}
	}

	client := pb.NewIPsecClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data, err := client.IPsecStats(ctx, &pb.IPsecStatsReq{})
	if err != nil {
		log.Println(err)
	}
	log.Println(data)
	defer disconnectConnection()
	return nil
}

// TestIpsec runs few basic tests establishing ipsec tunnels, version and stats
func TestIpsec(ctx context.Context, conn grpc.ClientConnInterface) {
	// IPsec
	c1 := pb.NewIPsecClient(conn)

	// Load IPsec connection
	localIpsec := pb.IPsecLoadConnReq{
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
			LocalAuth:  &pb.LocalAuth{Auth: pb.AuthType_PSK, Id: "hacker@strongswan.org"},
			RemoteAuth: &pb.RemoteAuth{Auth: pb.AuthType_PSK, Id: "server.strongswan.org"},
			Children: []*pb.Child{
				{
					Name: "opi-child",
					EspProposals: &pb.Proposals{
						CryptoAlg: []pb.CryptoAlgorithm{pb.CryptoAlgorithm_AES256GCM128},
						IntegAlg:  []pb.IntegAlgorithm{pb.IntegAlgorithm_SHA512},
						Dhgroups:  []pb.DiffieHellmanGroups{pb.DiffieHellmanGroups_CURVE25519},
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

	getVersion(ctx, c1)

	statsResp, err := c1.IPsecStats(ctx, &pb.IPsecStatsReq{})
	if err != nil {
		log.Fatalf("could not get IPsec stats")
	}
	log.Printf("IPsec stats\n%s", statsResp.GetStatus())

	rs1, err := c1.IPsecLoadConn(ctx, &localIpsec)
	if err != nil {
		log.Fatalf("could not load IPsec tunnel: %v", err)
	}
	log.Printf("Loaded: %v", rs1)

	// Bring the connection up
	initConn := pb.IPsecInitiateReq{
		Ike:   "opi-test",
		Child: "opi-child",
	}

	initRet, err := c1.IPsecInitiate(ctx, &initConn)
	if err != nil {
		log.Fatalf("could not initiate IPsec tunnel: %v", err)
	}
	log.Printf("Initiated: %v", initRet)

	// List the ikeSas
	ikeSas := pb.IPsecListSasReq{
		Ike: "opi-test",
	}

	listSasRet, err := c1.IPsecListSas(ctx, &ikeSas)
	if err != nil {
		log.Fatalf("could not list ikeSas: %v", err)
	}
	log.Printf("Returned ikeSas: %v", listSasRet)

	// List the connections
	listConn := pb.IPsecListConnsReq{
		Ike: "opi-test",
	}

	listConnsRet, err := c1.IPsecListConns(ctx, &listConn)
	if err != nil {
		log.Fatalf("could not list connections: %v", err)
	}
	log.Printf("Returned connections: %v", listConnsRet)

	// List the certificates
	listCerts := pb.IPsecListCertsReq{
		Type: "any",
	}

	listCertsRet, err := c1.IPsecListCerts(ctx, &listCerts)
	if err != nil {
		log.Fatalf("could not list certificates: %v", err)
	}
	log.Printf("Returned connections: %v", listCertsRet)

	// Ping across the tunnel.
	doPing()

	// Rekey the IKE_SA
	rekeyConn := pb.IPsecRekeyReq{
		Ike: "opi-test",
	}

	rekeyRet, err := c1.IPsecRekey(ctx, &rekeyConn)
	if err != nil {
		log.Fatalf("could not rekey IPsec tunnel: %v", err)
	}
	log.Printf("Rekeyed IKE_SA %s: %v", "opi-test", rekeyRet)

	doCleanup(ctx, c1)
}

func doCleanup(ctx context.Context, client pb.IPsecClient) {
	// Terminate the connection
	termConn := pb.IPsecTerminateReq{
		Ike: "opi-test",
	}

	termRet, err := client.IPsecTerminate(ctx, &termConn)
	if err != nil {
		log.Fatalf("could not terminate IPsec tunnel: %v", err)
	}
	log.Printf("Terminate: %v", termRet)

	// Unload
	unloadIpsec := pb.IPsecUnloadConnReq{
		Name: "opi-test",
	}

	rs2, err := client.IPsecUnloadConn(ctx, &unloadIpsec)
	if err != nil {
		log.Fatalf("could not unload IPsec tunnel: %v", err)
	}
	log.Printf("Unloaded: %v", rs2)
}

func getVersion(ctx context.Context, client pb.IPsecClient) {
	vresp, err := client.IPsecVersion(ctx, &pb.IPsecVersionReq{})
	if err != nil {
		log.Fatalf("could not get IPsec version")
	}
	log.Printf("Daemon  [%v]", vresp.GetDaemon())
	log.Printf("Version [%v]", vresp.GetVersion())
	log.Printf("Sysname [%v]", vresp.GetSysname())
	log.Printf("Release [%v]", vresp.GetRelease())
	log.Printf("Machine [%v]", vresp.GetMachine())
}

func doPing() {
	// .NOTE: The container this test runs in is linked to the appropriate
	//        strongSwan container.
	pinger, err := ping.NewPinger(pingaddr)
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

func dialConnection() error {
	var err error
	conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	} else {
		log.Println("GRPC connection closed successfully")
	}
}
