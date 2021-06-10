package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <hostname>\n", os.Args[0])
		os.Exit(1)
	}

	hostname := os.Args[1]

	ipAddr, isErr := net.ResolveIPAddr("ip", hostname)

	if isErr != nil {
		fmt.Println("Resolution Error:", isErr.Error())
		os.Exit(1)
	}

	fmt.Println("Resolved Address:", ipAddr.String())

	os.Exit(0)
}
