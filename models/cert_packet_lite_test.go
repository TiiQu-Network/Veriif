package models

import (
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var(
	liteJSON = `{"d":{"cert_name":"Biomedical Technology","date":"20th July 2032","email":"samuel.hr@tiiqu.com","expires_on":"nil","level":"SECOND CLASS HONOURS (upper division)","recipient_name":"Samuel Hawksby-Robinson","type":"MASTER of ENGINEERING"},"h":"g4VmucMWywSZvUeLv5phJXHLGp016iyVowMXfRULPeE=","p":["r-Q4rYkJmwpq+Mxb1sG/pgiTopqIz7fjAM4K8S/IEB8QY=","r-Y0KQjKUFodAHPil8V7y5KAA1oIeKneH8y23G0x7Owmo="],"r":"vTzswCve6CtRoh42/ue2105BxM9w0dMygpJAuy5t1Bo="}`
)

func TestCertPacketLite_ToCertPacket(t *testing.T) {
	var cpl CertPacketLite
	err := json.Unmarshal([]byte(liteJSON), &cpl)
	if err != nil {
		t.Error(err)
	}

	spew.Dump(cpl)

	cp, err := cpl.ToCertPacket()
	if err != nil {
		t.Error(err)
	}

	spew.Dump(cp)
}
