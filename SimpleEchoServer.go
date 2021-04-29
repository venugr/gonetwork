package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	service := ":1201"

	tcpAddr, isErr := net.ResolveTCPAddr("tcp", service)
	checkError(isErr)

	listener, isErr := net.ListenTCP("tcp", tcpAddr)
	checkError(isErr)

	for {

		conn, isErr := listener.Accept()
		if isErr != nil {
			continue
		}

		handleConnection(conn)
		conn.Close()
	}
}

func handleConnection(conn net.Conn) {
	var buf [512]byte

	for {
		n, isErr := conn.Read(buf[0:])
		if isErr != nil {
			return
		}

		fmt.Println(string(buf[0:]))

		_, isErr = conn.Write(buf[0:n])
		if isErr != nil {
			return
		}

		return
	}
}

func checkError(isErr error) {
	if isErr != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", isErr.Error())
		os.Exit(1)
	}

}
