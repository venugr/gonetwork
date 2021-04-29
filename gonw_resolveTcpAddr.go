package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <network-type> <hostname:port>\n", os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]
	hostname := os.Args[2]

	tcpAddr, isErr := net.ResolveTCPAddr(networkType, hostname)

	if isErr != nil {
		fmt.Println("Error:", isErr.Error())
		os.Exit(2)
	}

	fmt.Println("IP:", tcpAddr.IP.String(), "Port:", tcpAddr.Port)
	os.Exit(0)

}
