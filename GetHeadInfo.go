package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <hostname:port> <echo-text>\n", os.Args[0])
		os.Exit(1)
	}

	hostname := os.Args[1]
	echoText := os.Args[2]

	tcpAddr, isErr := net.ResolveTCPAddr("tcp4", hostname)
	checkError(isErr)

	conn, isErr := net.DialTCP("tcp", nil, tcpAddr)
	checkError(isErr)

	//_, isErr = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	_, isErr = conn.Write([]byte(echoText))
	checkError(isErr)
	result, isErr := ioutil.ReadAll(conn)
	checkError(isErr)
	conn.Close()
	fmt.Println(string(result))
}

func checkError(isErr error) {
	if isErr != nil {
		fmt.Println("Error:", isErr.Error())
		os.Exit(2)
	}
}
