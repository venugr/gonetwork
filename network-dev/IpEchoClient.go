package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <hostname:port> <text>\n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	text := os.Args[2]

	conn, isErr := net.Dial("tcp", service)
	checkError(isErr)

	_, isErr = conn.Write([]byte(text))
	checkError(isErr)

	result, isErr := ReadFully(conn)
	checkError(isErr)

	fmt.Println(string(result))
}

func ReadFully(conn net.Conn) ([]byte, error) {

	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		n, isErr := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if isErr != nil {
			if isErr == io.EOF {
				break
			}
			return nil, isErr
		}
	}

	return result.Bytes(), nil

}

func checkError(isErr error) {
	if isErr != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", isErr.Error())
		os.Exit(1)
	}

}
