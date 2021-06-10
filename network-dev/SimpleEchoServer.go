package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
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

		rand.Seed(time.Now().UnixNano())

		trand := rand.Intn(15-1) + 1
		handleConnection(conn, trand)
		conn.Close()
	}
}

func handleConnection(conn net.Conn, trand int) {
	var buf [512]byte

	for {
		n, isErr := conn.Read(buf[0:])
		if isErr != nil {
			return
		}

		fmt.Println(string(buf[0:]))

		fmt.Println("wait for", trand, "seconds...")
		time.Sleep(time.Duration(trand) * time.Second)

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
