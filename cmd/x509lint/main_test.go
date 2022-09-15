package main

import (
	"log"
	"testing"
)

func TestDoLint(t *testing.T) {
	result := doLint("/Users/microshine/github/pv/x509-linter/cert_no_serial.pem")
	log.Println(result)
}
