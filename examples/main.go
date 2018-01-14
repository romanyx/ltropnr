package main

import (
	"bytes"
	"log"

	"github.com/romanyx/ltropnr"
)

const msg = `Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon

Message body
`

func main() {
	opnr := ltropnr.New()
	if err := opnr.Send(bytes.NewBufferString(msg)); err != nil {
		log.Fatalf("open letter in browser failed: %s", err)
	}
}
