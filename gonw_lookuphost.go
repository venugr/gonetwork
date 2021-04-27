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

	addrs, isErr := net.LookupHost(hostname)

	if isErr != nil {
		fmt.Println("Error:", isErr.Error())
		os.Exit(2)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}

	cname, isErr := net.LookupCNAME(hostname)

	if isErr != nil {
		fmt.Println("Error:", isErr.Error())
		os.Exit(3)
	}

	fmt.Println("Cname:", cname)
	os.Exit(0)
}
