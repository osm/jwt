package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// Define a new Parser type
type Parser int

// Parse will parse the given token and return the header and claims
func (*Parser) Parse(token string) (string, string, error) {
	// The different parts should be separated with a dot
	p := strings.Split(token, ".")
	if len(p) < 2 {
		return "", "", fmt.Errorf("error: the token doesn't look right")
	}

	// We are only interested in the header and claims
	var r []string
	for i := 0; i < 2; i++ {
		// Decode the data
		d, err := base64.StdEncoding.DecodeString(p[i])
		if err != nil {
			return "", "", err
		}

		// Run the decoded data through json.Indent to make it nice
		var o bytes.Buffer
		err = json.Indent(&o, d, "", "\t")
		if err != nil {
			return "", "", err
		}

		// Append the indented token to the slice
		r = append(r, o.String())
	}

	// Return it
	return r[0], r[1], nil
}
