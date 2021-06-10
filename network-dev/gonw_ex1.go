package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <ip-addr>\n", os.Args[0])
		os.Exit(1)
	}

	argAddr := os.Args[1]

	ipAddr := net.ParseIP(argAddr)

	if ipAddr == nil {
		fmt.Println("inavlid address")
	} else {
		fmt.Println("The address is:", ipAddr.String())
	}

	os.Exit(0)
}
