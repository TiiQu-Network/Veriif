// TODO create a separate package just for these structs, then import them in Veriif and CerTiiF.
package models

type CertPacket struct {
	Data        CertData      `json:"data"`
	Hash        EncodingGroup `json:"hash"`
	MerkleProof MerkleProof   `json:"merkle_proof"`
	MerkleRoot  EncodingGroup `json:"merkle_root"`
	KeyType     string        `json:"key_type"`
	PublicKey   string        `json:"public_key"`
	Signature   EncodingGroup `json:"signature"`
}

type CertData struct {
	RecipientName string `json:"name" csv:"name"`
	CertType      string `json:"cert_type" csv:"type"`
	CertLevel     string `json:"cert_level" csv:"level"`
	CertName      string `json:"cert_name" csv:"cert_name"`
	CertDate      string `json:"cert_date" csv:"date"`
}

type MerkleProof []map[string]EncodingGroup
