// TODO create a separate package just for these structs, then import them in Veriif and CerTiiF.
package models

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strconv"
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
	var data [][]byte

	if len(e.Hex) > 0 {
		h, _ := hex.DecodeString(e.Hex)
		data = append(data, h)
	}

	if len(e.Base64) > 0 {
		b, _ := base64.StdEncoding.DecodeString(e.Base64)
		data = append(data, b)
	}

	for i, b := range data {
		if !bytes.Equal(b, data[0]) {
			return []byte{}, errors.New("decode mismatch between " + strconv.Itoa(i) + " and 0")
		}
	}

	return data[0], nil
}
