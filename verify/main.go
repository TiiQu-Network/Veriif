package verify

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/big"
	"regexp"
	"time"

	"github.com/onrik/gomerkle"
	"github.com/pkg/errors"

	"github.com/TiiQu-Network/Veriif/messages"
	"github.com/TiiQu-Network/Veriif/models"
	"github.com/TiiQu-Network/Veriif/ethereum"
)

var (
	expiresOn = "expires_on"

	mandatoryFields = []string{
		expiresOn,
	}
)

func Verify(data models.CertPacket, e ethereum.Caller) (map[string]bool, error) {
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
		output["check_hash"] = false
		return output, err
	}
	output["check_hash"] = true

	// Check that the "expires_on" is not expired
	err = checkHasExpired(data.Data[expiresOn])
	if err != nil {
		output["check_expired"] = false
		return output, err
	}
	output["check_expired"] = true

	// check public key is known, if the key is set
	err = checkPublicKeyRegistered(data, e)
	if err != nil {
		output["check_pk_registered"] = false
		return output, err
	}
	output["check_pk_registered"] = true

	// Check the signature verifies, if the key is set
	err = checkSigVerifies(certHash, data)
	if err != nil {
		output["check_sig"] = false
		return output, err
	}
	output["check_sig"] = true

	// Check the Merkle proof verifies
	mr, err := checkMerkleProof(certHash, data)
	if err != nil {
		output["check_merkle_proof"] = false
		return output, err
	}
	output["check_merkle_proof"] = true

	// Check merkle root is on blockchain
	err = checkMerkelRoot(mr, e)
	if err != nil {
		output["check_merkle_root"] = false
		return output, err
	}
	output["check_merkle_root"] = true

	// check hash has been suspended
	err = checkHashSuspended(certHash, e)
	if err != nil {
		output["check_suspended"] = false
		return output, err
	}
	output["check_suspended"] = true

	// Check hash or merkle root has been revoked
	err = checkHashRevoked(certHash, mr, e)
	if err != nil {
		output["check_revocation"] = false
		return output, err
	}
	output["check_revocation"] = true

	return output, nil
}

func ExtractCert(content []byte) (models.CertPacket, error) {
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
	errorHeader := "Hash"

	certHash, err := data.Hash.Decode()
	if err != nil {
		return nil, errors.Wrap(err, errorHeader)
	}

	if bytes.Equal(calcHash[:], certHash) {
		messages.PrintSuccess("Hash match")
	} else {
		return nil, errors.New(errorHeader + ": does not match")
	}

	return certHash, nil
}

func checkHasExpired(expireDate string) error {
	errorHeader := "Expiry date"

	if expireDate != "nil" {

		// Check date is in the correct format
		t, err := time.Parse("2006-01-02", expireDate)
		if err != nil {
			return errors.New(fmt.Sprintf(
				"%s: '%s' is not a valid date format. Please format in date in YYYY-MM-DD.",
				errorHeader,
				expireDate),
			)
		}

		// Check date is in the future
		if t.Before(time.Now()) {
			return errors.New(fmt.Sprintf(
				"%s: %s must be a date in the future. %s is not a valid date.",
				errorHeader,
				expireDate,
				expireDate),
			)
		}
	}

	return nil
}

func checkSigVerifies(certHash []byte, data models.CertPacket) error {
	errorHeader := "Signature"

	switch data.KeyType {
	case "rsa":
		pk, err := parsePublicKey([]byte(data.PublicKey))
		if err != nil {
			return err
		}

		sig, err := data.Signature.Decode()
		if err != nil {
			return errors.Wrap(err, errorHeader)
		}

		err = rsa.VerifyPKCS1v15(pk, crypto.SHA256, certHash, sig)
		if err != nil {
			return errors.Wrap(err, errorHeader+": does not verify")
		} else {
			messages.PrintSuccess(errorHeader + ": verification")
		}

		break
	case "ecdsa":
		// Parse and decode the signature
		sig, err := data.Signature.Decode()
		if err != nil {
			return errors.Wrap(err, errorHeader)
		}

		var eSig struct {
			R, S *big.Int
		}
		asn1.Unmarshal(sig, &eSig)

		// Parse and decode the public key
		block, _ := pem.Decode([]byte(data.PublicKey))
		pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return errors.Wrap(err, errorHeader)
		}

		if !ecdsa.Verify(pubKey.(*ecdsa.PublicKey), certHash, eSig.R, eSig.S) {
			return errors.New(errorHeader + ": does not verify")
		}
	default:
		return errors.New(errorHeader + ": unknown key type - " + data.KeyType)
	}

	return nil
}

func checkMerkleProof(certHash []byte, data models.CertPacket) ([]byte, error) {
	errorHeader := "Merkle proof"

	mp, err := proofFromModel(data.MerkleProof)
	if err != nil {
		return nil, err
	}

	mr, err := data.MerkleRoot.Decode()
	if err != nil {
		return nil, errors.Wrap(err, "Merkle root")
	}

	mt := gomerkle.NewTree(sha256.New())
	if mt.VerifyProof(mp, mr, certHash) {
		messages.PrintSuccess(errorHeader + ": verification")
	} else {
		return nil, errors.New(errorHeader + ": does not verify")
	}

	return mr, nil
}

func checkMerkelRoot(merkleRoot []byte, e ethereum.Caller) error {
	errorHeader := "Merkle proof"

	exists, err := rootExistsOnChain(merkleRoot, e)
	if err != nil {
		return err
	}

	if exists {
		messages.PrintSuccess(errorHeader + ": on chain")
	} else {
		return errors.New(errorHeader + ": not on chain")
	}

	return nil
}

// parsePublicKey parses a PEM encoded public key.
func parsePublicKey(pemBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("no valid key found")
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

func checkPublicKeyRegistered(data models.CertPacket, e ethereum.Caller) error {
	errorHeader := "Public key"

	pkb := []byte(data.PublicKey)
	pb, _ := pem.Decode(pkb)
	if pb == nil {
		return errors.New(errorHeader + ": no key found")
	}

	hash := sha256.Sum256(pb.Bytes)
	registered, err := e.CertRegContract.PublicKeyIsRegistered(e.CallOpts, hash)
	if err != nil {
		return err
	}

	if !registered {
		return errors.New(errorHeader + ": is not registered")
	}

	return nil
}

func proofFromModel(md models.MerkleProof) (gomerkle.Proof, error) {
	var pf gomerkle.Proof

	for _, item := range md {
		for dir, hashes := range item {
			b, err := hashes.Decode()
			if err != nil {
				return pf, errors.Wrap(err, "Merkle proof")
			}

			p := map[string][]byte{dir: b}
			pf = append(pf, p)
		}
	}

	return pf, nil
}

func rootExistsOnChain(root []byte, e ethereum.Caller) (bool, error) {
	var mr32 [32]byte
	copy(mr32[:], root)

	return e.CertRegContract.RootExists(e.CallOpts, mr32)
}

func checkHashSuspended(certHash []byte, e ethereum.Caller) error {
	errorHeader := "Suspended"

	var ch32 [32]byte
	copy(ch32[:], certHash)
	chr, err := e.CertRegContract.IsSuspended(e.CallOpts, ch32)
	if err != nil {
		return errors.Wrap(err, errorHeader)
	}
	if chr == true {
		return errors.New(errorHeader + ": certificate has been suspended")
	}

	return nil
}

func checkHashRevoked(certHash []byte, root []byte, e ethereum.Caller) error {
	errorHeader := "Revoked"

	var ch32 [32]byte
	copy(ch32[:], certHash)
	chr, err := e.CertRegContract.IsRevoked(e.CallOpts, ch32)
	if err != nil {
		return errors.Wrap(err, errorHeader)
	}
	if chr == true {
		return errors.New(errorHeader + ": certificate has been revoked")
	}

	var r32 [32]byte
	copy(r32[:], root)
	rr, err := e.CertRegContract.IsRevoked(e.CallOpts, r32)
	if err != nil {
		return errors.Wrap(err, errorHeader)
	}
	if rr == true {
		return errors.New(errorHeader + ": certificate batch has been revoked")
	}

	messages.PrintSuccess("Certificate not revoked")

	return nil
}
