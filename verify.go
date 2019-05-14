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
	"github.com/onrik/gomerkle"
	"github.com/pkg/errors"
	"regexp"
	"time"
)

var (
	expiresOn = "expires_on"

	mandatoryFields = []string{
		expiresOn,
	}
)

func verify(data models.CertPacket) (map[string]bool, error) {
	output := map[string]bool{}

	// Check that mandatory fields are present
	err := hasFields(mandatoryFields, data.Data, "missing mandatory field(s) : ")
	if err != nil {
		return nil, err
	}

	// Calculate the hash of the data
	calcHash, err := calculateDataHash(data)
	if err != nil {
		return nil, err
	}

	// Check the calculated hash matches the cert hash
	certHash, err := checkHashMatch(calcHash[:], data)
	if err != nil {
		return output, err
	}
	output["check_hash"] = true

	// Check that the "expires_on" is not expired
	err = checkHasExpired(data.Data[expiresOn])
	if err != nil {
		return output, err
	}
	output["check_expired"] = true

	// check public key is known
	err = checkPublicKeyRegistered(data)
	if err != nil {
		return output, err
	}
	output["check_pk_registered"] = true

	// Check the signature verifies
	err = checkSigVerifies(certHash, data)
	if err != nil {
		return output, err
	}
	output["check_sig"] = true

	// Check the Merkle proof is verifies
	mr, err := checkMerkleProof(certHash, data)
	if err != nil {
		return output, err
	}
	output["check_merkle_proof"] = true

	// Check merkle root is on blockchain
	err = checkMerkelRoot(mr)
	if err != nil {
		return output, err
	}
	output["check_merkle_root"] = true

	// check hash has been suspended
	err = checkHashSuspended(certHash)
	if err != nil {
		return output, err
	}
	output["check_suspended"] = true

	// Check hash or merkle root has been revoked
	err = checkHashRevoked(certHash, mr)
	if err != nil {
		return output, err
	}
	output["check_revocation"] = true

	return output, nil
}

func extractCert(content []byte) (models.CertPacket,error){
	certPacket := models.CertPacket{}

	// Get the cert data from a file
	cert, err := findCert(content)
	if err != nil {
		return certPacket, err
	}

	// Parse the cert data into the expected struct
	err = json.Unmarshal(cert, &certPacket)
	if err != nil {
		return certPacket, err
	}

	return certPacket, nil
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

func hasFields(fields []string, data map[string]string, errMsg string) error {
	// Set up the mandatory field check
	result := map[string]bool{}
	for _, mandatory := range fields {
		result[mandatory] = false
	}

	// Check the fields
	for key := range data {
		for _, mandatory := range fields {
			if key == mandatory {
				result[mandatory] = true
			}
		}
	}

	// Check for fails
	message := ""
	for man, res := range result {
		if !res {
			if message == "" {
				message += errMsg + man
			} else {
				message += ", " + man
			}
		}
	}
	if message == "" {
		return nil
	} else {
		return errors.New(message)
	}
}

func calculateDataHash(cert models.CertPacket) ([32]byte, error) {
	jn, err := json.Marshal(cert.Data)
	if err != nil {
		return [32]byte{}, err
	}

	hash := sha256.Sum256(jn)
	return hash, nil
}

func checkHashMatch(calcHash []byte, data models.CertPacket) ([]byte, error) {
	certHash, err := data.Hash.Decode()
	if err != nil {
		return nil, err
	}

	if bytes.Equal(calcHash[:], certHash) {
		printSuccess("Hash match")
	} else {
		return nil, errors.New("Hash does not match")
	}

	return certHash, nil
}

func checkHasExpired(expireDate string) error {
	if expireDate != "nil" {

		// Check date is in the correct format
		t, err := time.Parse("2006-01-02", expireDate)
		if err != nil {
			return errors.New(fmt.Sprintf("'%s' is not a valid date format. Please format in date in YYYY-MM-DD.", expireDate))
		}

		// Check date is in the future
		if t.Before(time.Now()) {
			return errors.New(fmt.Sprintf(expireDate+" must be a date in the future. %s is not a valid date.", expireDate))
		}
	}

	return nil
}

func checkSigVerifies(certHash []byte, data models.CertPacket) error {
	pk, err := parsePublicKey([]byte(data.PublicKey))
	if err != nil {
		return err
	}

	sig, err := data.Signature.Decode()
	if err != nil {
		return err
	}

	err = rsa.VerifyPKCS1v15(pk, crypto.SHA256, certHash, sig)
	if err != nil {
		return errors.Wrap(err, "Signature does not verify")
	} else {
		printSuccess("Signature verification")
	}

	return nil
}

func checkMerkleProof(certHash []byte, data models.CertPacket) ([]byte, error) {
	mp, err := proofFromModel(data.MerkleProof)
	if err != nil {
		return nil, err
	}

	mr, err := data.MerkleRoot.Decode()
	if err != nil {
		return nil, err
	}

	mt := gomerkle.NewTree(sha256.New())
	if mt.VerifyProof(mp, mr, certHash) {
		printSuccess("Merkle proof verification")
	} else {
		return nil, errors.New("Merkle proof does not verify")
	}

	return mr, nil
}

func checkMerkelRoot(merkleRoot []byte) error {
	exists, err := rootExistsOnChain(merkleRoot)
	if err != nil {
		return err
	}

	if exists {
		printSuccess("Merkle root on chain")
	} else {
		return errors.New("Merkle root not on chain")
	}

	return nil
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
		return nil, errors.Errorf("unsupported key type %q", block.Type)
	}
}

func checkPublicKeyRegistered(data models.CertPacket) error {
	pkb := []byte(data.PublicKey)
	pb, r := pem.Decode(pkb)
	if pb == nil {
		return errors.New("no key found")
	}
	spew.Dump(pb, r)

	hash := sha256.Sum256(pb.Bytes)
	registered, err := certRegContract.PublicKeyIsRegistered(callOpts, hash)
	if err != nil {
		return err
	}

	if !registered {
		return errors.New("Public key if not registered")
	}

	return nil
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

func rootExistsOnChain(root []byte) (bool, error) {
	var mr32 [32]byte
	copy(mr32[:], root)

	return certRegContract.RootExists(callOpts, mr32)
}

func checkHashSuspended(certHash []byte) error {

	var ch32 [32]byte
	copy(ch32[:], certHash)
	chr, err := certRegContract.IsSuspended(callOpts, ch32)
	if err != nil {
		return err
	}
	if chr == true {
		return errors.New("Certificate has been suspended")
	}

	return nil
}

func checkHashRevoked(certHash []byte, root []byte) error {
	var ch32 [32]byte
	copy(ch32[:], certHash)
	chr, err := certRegContract.IsRevoked(callOpts, ch32)
	if err != nil {
		return err
	}
	if chr == true {
		return errors.New("Certificate has been revoked")
	}

	var r32 [32]byte
	copy(r32[:], root)
	rr, err := certRegContract.IsRevoked(callOpts, r32)
	if err != nil {
		return err
	}
	if rr == true {
		return errors.New("Certificate batch has been revoked")
	}

	printSuccess("Certificate not revoked")

	return nil
}
