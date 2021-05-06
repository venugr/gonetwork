package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <:port>")
		os.Exit(1)
	}

	service := os.Args[1]

	listener, isErr := net.Listen("tcp", service)
	checkError(isErr)

	for {
		conn, isErr := listener.Accept()
		if isErr != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	defer conn.Close()

	var buf [512]byte

	for {
		n, isErr := conn.Read(buf[0:])
		if isErr != nil {
			return
		}

		ip := conn.RemoteAddr().String()
		rStr := ip + " => " + string(buf[0:n])
		_, isErr = conn.Write([]byte(rStr))

		fmt.Println(string(buf[0:n]))
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
