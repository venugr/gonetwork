package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <network-type> <service-name>\n", os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]
	serviceName := os.Args[2]

	portNum, isErr := net.LookupPort(networkType, serviceName)

	if isErr != nil {
		fmt.Println("Error:", isErr.Error())
		os.Exit(2)
	}

	fmt.Println("Service Port#", portNum)
	os.Exit(0)
}
