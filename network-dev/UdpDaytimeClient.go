package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <hostname:port>\n", os.Args[0])
		os.Exit(1)
	}

	hostName := os.Args[1]

	udpAddr, isErr := net.ResolveUDPAddr("udp", hostName)
	checkError(isErr)

	udpConn, isErr := net.DialUDP("udp", nil, udpAddr)
	checkError(isErr)

	_, isErr = udpConn.Write([]byte("anything"))
	checkError(isErr)

	var buf [512]byte

	n, isErr := udpConn.Read(buf[0:])
	checkError(isErr)

	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}

func checkError(isErr error) {
	if isErr != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", isErr.Error())
		os.Exit(1)
	}

}
