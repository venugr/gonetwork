package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := ":1200"

	tcpAddr, isErr := net.ResolveTCPAddr("tcp", service)
	checkError(isErr)

	listener, isErr := net.ListenTCP("tcp", tcpAddr)
	checkError(isErr)

	for {
		conn, isErr := listener.Accept()
		if isErr != nil {
			continue
		}

		daytime := time.Now().String()

		conn.Write([]byte("DTS: " + daytime + "\n"))
		conn.Close()
	}
}

func checkError(isErr error) {
	if isErr != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", isErr.Error())
		os.Exit(1)
	}

}
