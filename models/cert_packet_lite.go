package models

import (
	"errors"
	"strings"
	"encoding/pem"
	"encoding/base64"
)

type CertPacketLite struct {
	D CertData `json:"d"` // Certificate [D]ata
	H string   `json:"h"` // [H]ash of certificate data
	P []string `json:"p"` // Merkle [P]roof
	R string   `json:"r"` // Merkle [R]oot
	T string   `json:"t"` // Signing key [T]ype
	K string   `json:"k"` // Signing [K]ey public element
	S string   `json:"s"` // Signing key [S]ignature of certificate hash
}

func (c CertPacketLite) ToCertPacket() (*CertPacket, error) {
	var out CertPacket

	mp, err := c.pToMerkleProof(c.P)
	if err != nil {
		return nil, err
	}

	out.Data = c.D
	out.Hash = c.stringToEncodingGroup(c.H)
	out.MerkleProof = *mp
	out.MerkleRoot = c.stringToEncodingGroup(c.R)
	out.KeyType = c.T

	out.PublicKey, err = c.kToPEMFormat()
	if err != nil {
		return nil, err
	}

	out.Signature = c.stringToEncodingGroup(c.S)

	return &out, nil
}

func (c CertPacketLite) stringToEncodingGroup(encoding string) EncodingGroup {
	return EncodingGroup{Base64:encoding}
}

func(c CertPacketLite) pToMerkleProof(proofs []string) (*MerkleProof, error) {
	var out MerkleProof

	for _, p := range proofs {

		// Lite proofs are formatted `r-Y0KQjKUFodAHPil8V7y5KAA1oIeKneH8y23G0x7Owmo=`
		// First step split the string into the position and the base64 hash
		es := strings.Split(p, "-")
		if len(es) != 2 {
			return nil, errors.New("proof format unknown: " + p)
		}

		// Convert the position character into the recognised position word. "right" or "left"
		var pos string
		switch es[0] {
		case "r":
			pos = "right"
		case "l":
			pos = "left"
		default:
			return nil, errors.New("Unknown position " + es[0])
		}

		// Create Merkle proof row
		row := map[string]EncodingGroup{
			pos: c.stringToEncodingGroup(es[1]),
		}

		out = append(out, row)
	}

	return &out, nil
}

func (c CertPacketLite) kToPEMFormat() (string, error) {
	// Decode the base64 public key string to bytes
	dpk, err := base64.StdEncoding.DecodeString(c.K)
	if err != nil {
		return "", err
	}

	// Parse the public key bytes to PEM format
	pemkey := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: dpk,
	}

	// Encode the PEM block into a string and return
	return string(pem.EncodeToMemory(pemkey)), nil
}
