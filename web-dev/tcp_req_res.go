package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	request(conn)
	response(conn)

}

func request(conn net.Conn) {

	ok := true
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ok {
			fmt.Printf("*** METHOD: \x1b[37;1m%s\x1b[0m ***\n", strings.Fields(ln)[0])
			ok = false
		}

		if ln == "" {
			break
		}
	}
}

func response(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong><h3>Hello World</h3></strong></body></html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
