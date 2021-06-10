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
	checkError1(isErr)

	listener, isErr := net.ListenTCP("tcp", tcpAddr)
	checkError1(isErr)

	for {

		conn, isErr := listener.Accept()
		if isErr != nil {
			continue
		}

		rand.Seed(time.Now().UnixNano())

		trand := rand.Intn(15-1) + 1
		go handleConnection1(conn, trand)

		// go handleConnection1(conn)
		// conn.Close()
	}
}

func handleConnection1(conn net.Conn, trand int) {
	var buf [512]byte

	defer conn.Close()

	for {
		n, isErr := conn.Read(buf[0:])
		if isErr != nil {
			return
		}

		text := string(buf[0:n])
		cmpText := "kill-your-self"

		fmt.Println(text)
		// fmt.Printf("%d", strings.Compare(text, cmpText))
		// fmt.Printf("Len(%s)=%d", text, len(text))
		// fmt.Printf("Len(%s)=%d", cmpText, len(text))

		if text == cmpText {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		fmt.Println("wait for", trand, "seconds...")
		time.Sleep(time.Duration(trand) * time.Second)

		_, isErr = conn.Write(buf[0:n])
		if isErr != nil {
			return
		}

		return
	}
}

func checkError1(isErr error) {
	if isErr != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", isErr.Error())
		os.Exit(1)
	}

}
