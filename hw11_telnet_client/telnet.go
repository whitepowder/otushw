package main

import (
	"io"
	"time"
)

type TelnetClient interface {
	// Place your code here
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	// Place your code here
	return nil
}

// Place your code here
// P.S. Author's solution takes no more than 50 lines
