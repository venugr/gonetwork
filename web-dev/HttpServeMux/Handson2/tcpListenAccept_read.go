package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panicln(err)
			continue
		}

		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	ok := true
	var rMethod, rURI string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ok {
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			ok = false
		}

		if ln == "" {
			break
		}
	}

	fmt.Println("CODE - Reached HERE!")

	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	body += "\n"
	body += "Method: " + rMethod
	body += "\n"
	body += "URI: " + rURI

	body = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
		</body>
		</html>
	`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	// fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
