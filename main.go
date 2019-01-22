package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
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
	spew.Dump(cert)

	// Parse the cert data into the expected struct
	data := models.CertPacket{}
	jn := json.Unmarshal(cert, &data)
	spew.Dump(jn, data)

	// Calculate the hash of the data
	calcHash, _ := calculateDataHash(data)

	// Check the calculated hash matches the cert hash
	certHash, _ := data.Hash.Decode()
	if bytes.Equal(calcHash[:], certHash) {
		spew.Dump("All good, hashes match")
	}

	// TODO check public key is known
	// TODO check the signature verifies
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
