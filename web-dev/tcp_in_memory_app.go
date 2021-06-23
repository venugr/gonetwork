package main

import (
	"bufio"
	"fmt"
	"io"
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

	io.WriteString(conn,
		"\r\nIN-MEMORY DATABASE\r\n\r\n"+
			"USE:\r\n"+
			"\tSET key value \r\n"+
			"\tGET key \r\n"+
			"\tDEL key \r\n\r\n"+
			"EXAMPLE:\r\n"+
			"\tSET fav chocolate \r\n"+
			"\tGET fav \r\n\r\n\r\n")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		case "GET":
			k := fs[1]
			fmt.Fprintf(conn, "%s\r\n", data[k])

		case "SET":
			if len(fs) != 3 {
				fmt.Fprintf(conn, "Expected a Value\r\n")
				continue
			}
			data[fs[1]] = fs[2]

		case "DEL":
			delete(data, fs[1])

		default:
			fmt.Fprintf(conn, "INVALID COMMAND: \x1b[31;1m%s\x1b[0m\r\n", fs[0])
		}

		fmt.Println(ln)
	}

}
