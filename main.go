package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Wait for data on stdin
	d, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Run the data through the parser
	var p Parser
	h, c, err := p.Parse(string(d))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Print it
	fmt.Printf("header:\n%s\nclaims:\n%s\n", h, c)
}
