package models

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var(
	liteJSON = `{"d":{"cert_name":"Biomedical Technology","date":"20th July 2032","email":"samuel.hr@tiiqu.com","expires_on":"nil","level":"SECOND CLASS HONOURS (upper division)","recipient_name":"Samuel Hawksby-Robinson","type":"MASTER of ENGINEERING"},"h":"g4VmucMWywSZvUeLv5phJXHLGp016iyVowMXfRULPeE=","p":["r-Q4rYkJmwpq+Mxb1sG/pgiTopqIz7fjAM4K8S/IEB8QY=","r-Y0KQjKUFodAHPil8V7y5KAA1oIeKneH8y23G0x7Owmo="],"r":"vTzswCve6CtRoh42/ue2105BxM9w0dMygpJAuy5t1Bo="}`
	queryStr = "D=city%3ATorino%0Acourse%3Aweb+development+course%0Adate%3A20th+July+2017%0Aduration%3A3-months%0Aemail%3Asamuel.hr%40tiiqu.com%0Arecipient_name%3ASamuel+Hawksby-Robinson%0Atitle%3APOWERCODER&H=68y58rrfjGICW6I1JiwVAZ0PDImNhbNXYpJs6gUrIFM%3D&K=MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE3VN8ylL%2FfdHTQpYHO7Se%2Bpt6RmvnUIPPNnj%2FA2VxKgaxLSVxvTyehEjLsGsO8co8n4FQWrZ0rvRtKLOWhaC6Zg%3D%3D&P=r-8sRlzjQrQODYg6F8hPVTlbaLaOX4DEXE1uwJ8sPTkl4%3D&P=r-e%2Fjt2JMCGFWy3dxIrzfMfsKSFxzyYLq8qBZgK%2BpLAsE%3D&R=bCo3yhwA7we%2Bsjv%2Fu%2Bjx6mulH5Vq%2FPUA2h0jnhTBclI%3D&S=MEUCIDErHa%2BPq6%2B6Noe25n%2FLvbXSAsqv3YLYe53nW42RQzvcAiEA5iHrZ%2BqCnEwayTipJ%2F8CLf6yIQXxoFsXIWcSqZUO9w4%3D&T=ecdsa"
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

func TestSchemeCertDataDecoder(t *testing.T) {
	var cpl CertPacketLite

	uv, err := url.ParseQuery(queryStr)
	if err != nil {
		t.Error(err)
	}

	err = cpl.UnmarshalSchemaQuery(uv)
	if err != nil {
		t.Error(err)
	}

	spew.Dump(cpl)
}
