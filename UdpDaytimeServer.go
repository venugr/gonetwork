package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {

	service := ":1200"

	udpAddr, isErr := net.ResolveUDPAddr("udp", service)
	checkError(isErr)

	udpConn, isErr := net.ListenUDP("udp", udpAddr)
	checkError(isErr)

	for {
		handleClient(udpConn)
	}

}

func handleClient(udpConn *net.UDPConn) {

	var buf [512]byte

	_, cliAddr, isErr := udpConn.ReadFromUDP(buf[0:])
	if isErr != nil {
		return
	}

	ip := cliAddr.IP.String()
	port := strconv.Itoa(cliAddr.Port)

	daytime := time.Now().String()
	udpConn.WriteToUDP([]byte(ip+":"+port+" => "+daytime), cliAddr)
}

func checkError(isErr error) {
	if isErr != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", isErr.Error())
		os.Exit(1)
	}

}
