package main

import "testing"

func TestCertTest(t *testing.T) {
	cert := []byte("asdkjasdkjabsdkhabsdkhbaskdbkasbdkakjsdkabskdba<certiif>I am a poodle</certiif>")
	expect := []byte("I am a poodle")

	result, err := findCert(cert)
	if err != nil {
		t.Error(err)
	}

	if string(expect) != string(result) {
		t.Error("does not match")
	}
}