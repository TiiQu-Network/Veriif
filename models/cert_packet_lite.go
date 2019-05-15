package models

import (
	"errors"
	"strings"
)

type CertPacketLite struct {
	D CertData `json:"d"`
	H string   `json:"h"`
	P []string `json:"p"`
	R string   `json:"r"`
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
