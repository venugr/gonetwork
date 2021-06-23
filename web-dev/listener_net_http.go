package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	ClrBlack   = "\x1b[30;1m"
	ClrRed     = "\x1b[31;1m"
	ClrGreen   = "\x1b[32;1m"
	ClrYellow  = "\x1b[33;1m"
	ClrBlue    = "\x1b[34;1m"
	ClrMagenta = "\x1b[35;1m"
	ClrCyan    = "\x1b[36;1m"
	ClrWhite   = "\x1b[37;1m"
	ClrUnColor = "\x1b[0m"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {

		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}

}

func handle(conn net.Conn) {

	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()

		rt := rot13([]byte(ln))
		fmt.Println(ln)
		fmt.Fprintf(conn, "%s - \x1b[31;1m%s\x1b[0m\n", ln, rt)
	}

}

func rot13(bs []byte) []byte {
	var br = make([]byte, len(bs))

	// 97 - 122 a - z
	// 65 - 90  A - Z

	for i, v := range bs {

		if v >= 97 && v <= 122 {
			if v <= 109 {
				br[i] = v + 13
			} else {
				br[i] = v - 13
			}
		}

		if v >= 65 && v <= 90 {
			if v <= 77 {
				br[i] = v + 13
			} else {
				br[i] = v - 13
			}
		}

	}

	return br

}

func handle1(conn net.Conn) {

	defer conn.Close()
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: \x1b[31;1m%s\x1b[0m\n", ln)
	}

	fmt.Println("Code Got Reached\n")
}
