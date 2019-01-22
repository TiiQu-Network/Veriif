package main

import (
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
