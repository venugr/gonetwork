package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <ip-addr>\n", os.Args[0])
		os.Exit(1)
	}

	argAddr := os.Args[1]

	ipAddr := net.ParseIP(argAddr)

	if ipAddr == nil {
		fmt.Println("inavlid address")
	} else {
		fmt.Println("The address is:", ipAddr.String())
	}

	defMask := ipAddr.DefaultMask()
	network := ipAddr.Mask(defMask)

	ones, bits := defMask.Size()

	fmt.Println("Adress is", ipAddr.String(),
		", Default Mask(hex):", defMask.String(),
		", Network: ", network.String(),
		", Default mask length is ", bits,
		", Leading ones count is ", ones,
	)
	os.Exit(0)
}
