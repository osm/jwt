package main

import (
	"testing"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
const header = `{
	"alg": "HS256",
	"typ": "JWT"
}`
const claims = `{
	"sub": "1234567890",
	"name": "John Doe",
	"admin": true
}`

func TestParser(t *testing.T) {
	var p Parser
	h, c, err := p.Parse(token)

	if err != nil {
		t.Errorf("parse failed, %v", err)
	}

	if h == "" {
		t.Errorf("h should not be empty")
	}

	if header != h {
		t.Errorf("header is in an unexpected format")
	}

	if c == "" {
		t.Errorf("c should not be empty")
	}

	if claims != c {
		t.Errorf("claims is in an unexpected format")
	}

}
