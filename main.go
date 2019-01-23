package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/Samyoul/Veriif/contracts"
	"github.com/Samyoul/Veriif/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/onrik/gomerkle"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// TODO add proper error handling / logging
func main() {
	color.Set(color.FgGreen, color.Bold)
	println(`____   ____           .__.__  _____`)
	println(`\   \ /   /___________|__|__|/ ____\`)
	println(` \   Y   // __ \_  __ \  |  \   __\ `)
	println(`  \     /\  ___/|  | \/  |  ||  |`)
	println(`   \___/  \_____>__|  |__|__||__|`)
	println("")
	color.Unset()

	println("application - Veriif (c) 2019 TiiQu Ltd")
	println("version - 0.1.0")
	println("")

	// Get file name from input
	for {
		reader := bufio.NewReader(os.Stdin)
		print("Enter certificate filename: ")
		color.Set(color.FgHiYellow)
		filename, _ := reader.ReadString('\n')
		filename = strings.TrimSpace(filename)
		println("- Processing : " + filename)
		color.Unset()

		err := verify(filename)
		if err != nil {
			color.Set(color.FgRed)
			println("- FAIL    - " + err.Error())
			color.Unset()

			color.Set(color.FgRed, color.Bold)
			println("|--------------------------------------------------------------|")
			println("| {x}  NOT VERIFIED - The certificate failed verification  {x} |")
			println("|--------------------------------------------------------------|")
			println("")
			color.Unset()

		} else {
			color.Set(color.FgGreen, color.Bold)
			println("|--------------------------------------------------------------|")
			println("| {+}     VERIFIED - The certificate has been verified     {+} |")
			println("|--------------------------------------------------------------|")
			println("")
			color.Unset()
		}
	}
}

func verify(filename string) error {
	// Get the cert file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Get the cert data from a file
	cert, err := findCert(content)
	if err != nil {
		return err
	}

	// Parse the cert data into the expected struct
	data := models.CertPacket{}
	err = json.Unmarshal(cert, &data)
	if err != nil {
		return err
	}

	// Calculate the hash of the data
	calcHash, err := calculateDataHash(data)
	if err != nil {
		return err
	}

	// Check the calculated hash matches the cert hash
	certHash, err := checkHashMatch(calcHash[:], data)
	if err != nil {
		return err
	}

	// TODO check public key is known

	// Check the signature verifies
	err = checkSigVerifies(certHash, data)
	if err != nil {
		return err
	}

	// Check the Merkle proof is verifies
	mr, err := checkMerkleProof(certHash, data)
	if err != nil {
		return err
	}

	// Check merkle root is on blockchain
	err = checkMerkelRoot(mr)
	if err != nil {
		return err
	}

	return nil
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

func checkHashMatch(calcHash []byte, data models.CertPacket) ([]byte, error) {
	certHash, err := data.Hash.Decode()
	if err != nil {
		return nil, err
	}

	if bytes.Equal(calcHash[:], certHash) {
		color.Set(color.FgGreen)
		println("- SUCCESS - Hash match")
		color.Unset()
	} else {
		return nil, errors.New("Hash does not match")
	}

	return certHash, nil
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
		return errors.Wrap(err,"Signature does not verify")
	} else {
		color.Set(color.FgGreen)
		println("- SUCCESS - Signature verification")
		color.Unset()
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
		color.Set(color.FgGreen)
		println("- SUCCESS - Merkle proof verification")
		color.Unset()
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
		color.Set(color.FgGreen)
		println("- SUCCESS - Merkle root on chain")
		color.Unset()
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

func rootExistsOnChain(root []byte) (bool, error) {
	contractAddress := common.HexToAddress("0x1099a30581552418062f09f701d15558253daae9")
	node := "https://kovan.infura.io/v3/87abed49b20c49bd99f9bc2645023f34"

	client, err := ethclient.Dial(node)
	if err != nil {
		return false, errors.Wrap(err,"Failed to connect to the Ethereum client: %v")
	}

	certRegContract, err := contracts.NewCertRegistryCaller(contractAddress, client)
	if err != nil {
		return false, err
	}

	var mr32 [32]byte
	copy(mr32[:], root)

	callOpts := &bind.CallOpts{
		Pending: false,
		Context: context.TODO(),
	}

	return certRegContract.RootExists(callOpts, mr32)
}
