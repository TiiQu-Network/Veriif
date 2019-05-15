// TODO create a separate package just for these structs, then import them in Veriif and CerTiiF.
package models

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

type EncodingGroup struct {
	Hex    string `json:"hex"`
	Base64 string `json:"base64"`
}

func NewEncodingGroup(hash []byte) EncodingGroup {
	return EncodingGroup{
		Base64: base64.StdEncoding.EncodeToString(hash),
		Hex:    hex.EncodeToString(hash),
	}
}

func (e *EncodingGroup) Decode() ([]byte, error) {
	data := map[string][]byte{}

	if len(e.Hex) > 0 {
		h, _ := hex.DecodeString(e.Hex)
		data["hex"] = h
	}

	if len(e.Base64) > 0 {
		b, _ := base64.StdEncoding.DecodeString(e.Base64)
		data["base64"] = b
	}

	// Get the first element index in the map
	var fei string
	for i := range data {
		fei = i
		break
	}

	// Compare byte values
	for i, b := range data {
		if !bytes.Equal(b, data[fei]) {
			return []byte{}, errors.New("decode mismatch between " + i + " and " + fei)
		}
	}

	return data[fei], nil
}
