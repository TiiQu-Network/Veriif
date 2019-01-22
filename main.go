package main

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/Samyoul/Veriif/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onrik/gomerkle"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"regexp"
)

// TODO add proper error handling / logging
func main() {

	// Get the cert data from a file
	content, _ := ioutil.ReadFile("example/cert_0.pdf")
	cert, _ := findCert(content)

	// Parse the cert data into the expected struct
	data := models.CertPacket{}
	json.Unmarshal(cert, &data)

	// Calculate the hash of the data
	calcHash, _ := calculateDataHash(data)

	// Check the calculated hash matches the cert hash
	certHash, _ := data.Hash.Decode()
	if bytes.Equal(calcHash[:], certHash) {
		println("SUCCESS - Hash match")
	} else {
		println("FAIL - Hash match")
	}

	// TODO check public key is known

	// Check the signature verifies
	pk, _ := parsePublicKey([]byte(data.PublicKey))
	sig, _ := data.Signature.Decode()
	err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, certHash, sig)
	if err != nil {
		spew.Dump("FAIL - Signature verification", err)
	} else {
		println("SUCCESS - Signature verification")
	}

	// Check the Merkle proof is verifies
	mp, _ := proofFromModel(data.MerkleProof)
	mr, _ := data.MerkleRoot.Decode()
	mt := gomerkle.NewTree(sha256.New())
	if mt.VerifyProof(mp, mr, certHash) {
		println("SUCCESS - Merkle proof verify")
	} else {
		println("FAIL - Merkle proof verify")
	}

	// TODO check merkle root is on blockchain
	// TODO create golang ABI for contract
	//  https://github.com/Samyoul/TiiQu-Platform/blob/386826d546d6a5be44e5b4ac968d920c7ac7757e/README.md#building-go-contract-interfaces
	//  https://stackoverflow.com/questions/42520417/how-to-call-an-ethereum-contract-from-go
	contractAddress := common.HexToAddress("0x1099a30581552418062f09f701d15558253daae9")
	node := "https://kovan.infura.io/v3/3acc90c4079a4c3daccb59d50e893e14"
	client, err := ethclient.Dial(node)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	spew.Dump(contractAddress, client)

}

func findCert(location []byte) ([]byte, error) {
	r, err := regexp.Compile("(?s)<certiif>(.*)</certiif>")
	if err != nil {
		return nil, err
	}

	match := r.FindSubmatch(location)

	if len(match) > 0 {
		return match[1], nil
	}

	return nil, errors.New("No Certiif certificate found")
}

func calculateDataHash(cert models.CertPacket) ([32]byte, error) {
	jn, err := json.Marshal(cert.Data)
	if err != nil {
		return [32]byte{}, err
	}

	hash := sha256.Sum256(jn)
	return hash, nil
}

// parsePublicKey parses a PEM encoded private key.
func parsePublicKey(pemBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("no key found")
	}

	switch block.Type {
	case "PUBLIC KEY":
		pk, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return pk, nil
	default:
		return nil, fmt.Errorf("unsupported key type %q", block.Type)
	}
}

func proofFromModel(md models.MerkleProof) (gomerkle.Proof, error) {
	var pf gomerkle.Proof

	for _, item := range md {
		for dir, hashes := range item {
			b, err := hashes.Decode()
			if err != nil {
				return pf, err
			}

			p := map[string][]byte{dir: b}
			pf = append(pf, p)
		}
	}

	return pf, nil
}