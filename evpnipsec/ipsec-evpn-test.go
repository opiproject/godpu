// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package evpnipsec implements the go library for OPI to be used to establish networking
package evpnipsec

import (
	"context"
	"log"

	pb "github.com/opiproject/opi-evpn-bridge/pkg/ipsec/gen/go"
)

// AddSA adds a new SA
func (c IPSecEvpnClientImpl) AddSA(ctx context.Context, src string, dst string, spi uint32, proto int32, ifID uint32, reqid uint32, mode int32, intrface string, encAlg int32, encKey string,
	intAlg int32, intKey string, replayWindow uint32, tfc uint32, encap int32, esn int32, copyDf int32, copyEcn int32, copyDscp int32, initiator int32, inbound int32,
	update int32) (*pb.AddSAResp, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		log.Println("THREE")
		return nil, err
	}
	defer closer()

	client := c.getIPSecClient(conn)
	data, err := client.AddSA(ctx, &pb.AddSAReq{
		SaId: &pb.SAIdentifier{
			Src:   src,
			Dst:   dst,
			Spi:   spi,
			Proto: pb.IPSecProtocol(proto),
			IfId:  ifID,
		},

		SaData: &pb.AddSAReqData{
			Reqid:        reqid,
			Mode:         pb.IPSecMode(mode),
			Interface:    intrface,
			EncAlg:       pb.CryptoAlgorithm(encAlg),
			EncKey:       []byte(encKey),
			IntAlg:       pb.IntegAlgorithm(intAlg),
			IntKey:       []byte(intKey),
			ReplayWindow: replayWindow,
			Tfc:          tfc,
			Encap:        pb.Bool(encap),
			Esn:          pb.Bool(esn),
			CopyDf:       pb.Bool(copyDf),
			CopyEcn:      pb.Bool(copyEcn),
			CopyDscp:     pb.DSCPCopy(copyDscp),
			Initiator:    pb.Bool(initiator),
			Inbound:      pb.Bool(inbound),
			Update:       pb.Bool(update),
		}})
	if err != nil {
		log.Printf("error creating logical bridge: %s\n", err)
		log.Println("FOUR")
		return nil, err
	}

	return data, nil
}

// CreateLogicalBridge creates an Logical Bridge an OPI server
/*func (c IPSecEvpnClientImpl) AddSA(ctx context.Context, sareq *pb.AddSAReq) (*pb.AddSAResp, error) {

	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getIPSecClient(conn)
	data, err := client.AddSA(ctx, sareq)
	if err != nil {
		log.Printf("error creating logical bridge: %s\n", err)
		return nil, err
	}

	return data, nil
}*/

// DelSA deletes an SA an OPI server
func (c IPSecEvpnClientImpl) DelSA(ctx context.Context, src string, dst string, spi uint32, proto int32, ifID uint32) (*pb.DeleteSAResp, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getIPSecClient(conn)
	data, err := client.DeleteSA(ctx, &pb.DeleteSAReq{
		SaId: &pb.SAIdentifier{
			Src:   src,
			Dst:   dst,
			Spi:   spi,
			Proto: pb.IPSecProtocol(proto),
			IfId:  ifID,
		}})

	if err != nil {
		log.Printf("error creating logical bridge: %s\n", err)
		return nil, err
	}

	return data, nil
}

// CreateLogicalBridge creates an Logical Bridge an OPI server
/*func (c IPSecEvpnClientImpl) DelSA(ctx context.Context, sareq *pb.DeleteSAReq) (*pb.DeleteSAResp, error) {

	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getIPSecClient(conn)
	data, err := client.DeleteSA(ctx, sareq)
	if err != nil {
		log.Printf("error creating logical bridge: %s\n", err)
		return nil, err
	}

	return data, nil
}*/
