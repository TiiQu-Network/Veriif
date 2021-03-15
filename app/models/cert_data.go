// TODO create a separate package just for these structs, then import them in Veriif and CerTiiF.
package models

import (
	"reflect"
	"strings"
)

type CertPacket struct {
	Data        CertData      `json:"data"`
	Hash        EncodingGroup `json:"hash"`
	MerkleProof MerkleProof   `json:"merkle_proof"`
	MerkleRoot  EncodingGroup `json:"merkle_root"`
	KeyType     string        `json:"key_type"`
	PublicKey   string        `json:"public_key"`
	Signature   EncodingGroup `json:"signature"`
}

type CertData map[string]string

type MerkleProof []map[string]EncodingGroup

func SchemeCertDataDecoder(s string) reflect.Value {
	cd := CertData{}

	data := strings.Split(s, "\n")
	for _, v := range data {
		row := strings.SplitN(v, ":", 2)
		cd[row[0]] = row[1]
	}

	return reflect.ValueOf(cd)
}
