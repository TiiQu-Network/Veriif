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
	"github.com/pkg/errors"
	"io/ioutil"
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
		println("All good, hashes match")
	}

	// TODO check public key is known

	// Check the signature verifies
	pk, _ := parsePublicKey([]byte(data.PublicKey))
	sig, _ := data.Signature.Decode()
	err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, certHash, sig)
	if err != nil {
		spew.Dump("Signature verification fail!", err)
	} else {
		println("Signature verification SUCCESS.")
	}

	// TODO check the Merkle proof is verifies
	// TODO check merkle root is on blockchain

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