// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package evpnipsec implements the ipsec related CLI commands
package evpnipsec

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/opiproject/godpu/cmd/common"
	"github.com/opiproject/godpu/evpnipsec"
	pb "github.com/opiproject/opi-evpn-bridge/pkg/ipsec/gen/go"
	"github.com/spf13/cobra"
)

// AddSaCommand Add Sa Command
func AddSaCommand() *cobra.Command {
	var (
		src          string
		dst          string
		spi          uint32
		proto        int32
		ifID         uint32
		reqid        uint32
		mode         int32
		intrface     string
		encAlg       string
		encKey       string
		intAlg       string
		intKey       string
		replayWindow uint32
		tfc          uint32
		encap        int32
		esn          int32
		copyDf       int32
		copyEcn      int32
		copyDscp     int32
		initiator    int32
		inbound      int32
		update       int32
	)
	// Create the map of string to CryptoAlgorithm
	var EncAlgorithms = map[string]pb.CryptoAlgorithm{
		"rsvd":           pb.CryptoAlgorithm_ENCR_RSVD,
		"null":           pb.CryptoAlgorithm_ENCR_NULL,
		"aes_cbc":        pb.CryptoAlgorithm_ENCR_AES_CBC,
		"aes_ctr":        pb.CryptoAlgorithm_ENCR_AES_CTR,
		"aes_ccm_icv_8":  pb.CryptoAlgorithm_ENCR_AES_CCM_8,
		"aes_ccm_icv_12": pb.CryptoAlgorithm_ENCR_AES_CCM_12,
		"aes_ccm_icv_16": pb.CryptoAlgorithm_ENCR_AES_CCM_16,
		"aes_gcm_icv_8":  pb.CryptoAlgorithm_ENCR_AES_GCM_8,
		"aes_gcm_icv_12": pb.CryptoAlgorithm_ENCR_AES_GCM_12,
		"aes_gcm_icv_16": pb.CryptoAlgorithm_ENCR_AES_GCM_16,
		"aes_gmac":       pb.CryptoAlgorithm_ENCR_NULL_AUTH_AES_GMAC,
		"chacha_poly":    pb.CryptoAlgorithm_ENCR_CHACHA20_POLY1305,
	}
	var IntAlgorithms = map[string]pb.IntegAlgorithm{
		"sha1_96":  pb.IntegAlgorithm_AUTH_HMAC_SHA1_96,
		"xcbc_96":  pb.IntegAlgorithm_AUTH_AES_XCBC_96,
		"cmac_96":  pb.IntegAlgorithm_AUTH_AES_CMAC_96,
		"gmac_128": pb.IntegAlgorithm_AUTH_AES_128_GMAC,
		"gmac_192": pb.IntegAlgorithm_AUTH_AES_192_GMAC,
		"gmac_256": pb.IntegAlgorithm_AUTH_AES_256_GMAC,
		"sha2_128": pb.IntegAlgorithm_AUTH_HMAC_SHA2_256_128,
		"sha2_192": pb.IntegAlgorithm_AUTH_HMAC_SHA2_384_192,
		"sha2_256": pb.IntegAlgorithm_AUTH_HMAC_SHA2_512_256,
		"none":     pb.IntegAlgorithm_NONE,
	}

	var cmd = &cobra.Command{
		Use:     "add-sa",
		Aliases: []string{"c"},
		Short:   "add-sa functionality",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			IPSecEvpnClient, err := evpnipsec.NewIPSecClient(addr, tlsFiles)
			if err != nil {
				log.Printf("error Adding SA: %s\n", err)
				log.Println("ONE")
			}

			data, err := IPSecEvpnClient.AddSA(ctx,
				src, dst, spi, proto, ifID, reqid, mode, intrface, int32(EncAlgorithms[encAlg]), encKey, int32(IntAlgorithms[intAlg]), intKey,
				replayWindow, tfc, encap, esn, copyDf, copyEcn, copyDscp, initiator, inbound, update,
			)
			if err != nil {
				log.Printf("error error Adding SA: %s\n", err)
				log.Println("TWO")
			}
			fmt.Println("Add SA Req marshaled successfully:", data)
		},
	}

	cmd.Flags().StringVar(&src, "src", "", "Source address or hostname")
	cmd.Flags().StringVar(&dst, "dst", "", "Destination address or hostname")
	cmd.Flags().Uint32Var(&spi, "spi", 0, "SPI")
	cmd.Flags().Int32Var(&proto, "proto", 0, "Protocol (ESP/AH)")
	cmd.Flags().Uint32Var(&ifID, "if_id", 0, "Interface ID")
	cmd.Flags().Uint32Var(&reqid, "reqid", 0, "Reqid")
	cmd.Flags().Int32Var(&mode, "mode", 0, "Mode (tunnel, transport...)")
	cmd.Flags().StringVar(&intrface, "interface", "", "Network interface restricting policy")
	cmd.Flags().StringVar(&encAlg, "enc_alg", "aes_cbc", "rsvd, null, aes_cbc, aes_ctr, aes_ccm_icv_8, aes_ccm_icv_12, aes_ccm_icv_16, aes_gcm_icv_8, aes_gcm_icv_12, aes_gcm_icv_16, aes_gmac, chacha_poly")
	cmd.Flags().StringVar(&encKey, "enc_key", "", "Encryption key")
	cmd.Flags().StringVar(&intAlg, "int_alg", "none", "Integrity protection algorithm: sha1_96, xcbc_96, cmac_96, gmac_128, gmac_192, gmac_256, sha2_128, sha2_192, sha2_256, none")
	cmd.Flags().StringVar(&intKey, "int_key", "", "Integrity protection key")
	cmd.Flags().Uint32Var(&replayWindow, "replay_window", 0, "Anti-replay window size")
	cmd.Flags().Uint32Var(&tfc, "tfc", 0, "Traffic Flow Confidentiality padding")
	cmd.Flags().Int32Var(&encap, "encap", 0, "Enable UDP encapsulation for NAT traversal")
	cmd.Flags().Int32Var(&esn, "esn", 0, "Mark the SA should apply to packets after processing")
	cmd.Flags().Int32Var(&copyDf, "copy_df", 0, "Copy the DF bit to the outer IPv4 header in tunnel mode")
	cmd.Flags().Int32Var(&copyEcn, "copy_ecn", 0, "Copy the ECN header field to/from the outer header")
	cmd.Flags().Int32Var(&copyDscp, "copy_dscp", 0, "Copy the DSCP header field to/from the outer header")
	cmd.Flags().Int32Var(&initiator, "initiator", 0, "TRUE if initiator of the exchange creating the SA")
	cmd.Flags().Int32Var(&inbound, "inbound", 0, "TRUE if this is an inbound SA")
	cmd.Flags().Int32Var(&update, "update", 0, "TRUE if an SPI has already been allocated for this SA")

	if err := cmd.MarkFlagRequired("src"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("dst"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("spi"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("if_id"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	return cmd
}

// DelSaCommand tests the  del SA
func DelSaCommand() *cobra.Command {
	var (
		src   string
		dst   string
		spi   uint32
		proto int32
		ifID  uint32
	)

	var cmd = &cobra.Command{
		Use:     "Del-sa",
		Aliases: []string{"c"},
		Short:   "add-sa functionality",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			tlsFiles, err := c.Flags().GetString(common.TLSFiles)
			cobra.CheckErr(err)

			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			IPSecEvpnClient, err := evpnipsec.NewIPSecClient(addr, tlsFiles)
			if err != nil {
				log.Printf("error Deleting SA %s\n", err)
			}
			data, err := IPSecEvpnClient.DelSA(ctx, src, dst, spi, proto, ifID)
			if err != nil {
				log.Printf("error Deleting SA  %s\n", err)
			}
			fmt.Println("Deleting SA successfully:", data)
		},
	}

	cmd.Flags().StringVar(&src, "src", "", "Source address or hostname")
	cmd.Flags().StringVar(&dst, "dst", "", "Destination address or hostname")
	cmd.Flags().Uint32Var(&spi, "spi", 0, "SPI")
	cmd.Flags().Int32Var(&proto, "proto", 0, "Protocol (ESP/AH)")
	cmd.Flags().Uint32Var(&ifID, "if_id", 0, "Interface ID")
	if err := cmd.MarkFlagRequired("src"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("dst"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("spi"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("if_id"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	return cmd
}

// NewEvpnIPSecCommand tests the  inventory
func NewEvpnIPSecCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "evpnipsec",
		Aliases: []string{"g"},
		Short:   "Tests ipsec functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatalf("[ERROR] %s", err.Error())
			}
		},
	}

	cmd.AddCommand(AddSaCommand())
	cmd.AddCommand(DelSaCommand())
	return cmd
}
