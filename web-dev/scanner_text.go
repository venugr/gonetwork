package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	s := "This is text..\nwith multiple lines.\nThe scanner will scan the text and return lines code.\nTest the script."

	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("\n\n-----------------------------\n\n")

	scanner = bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("\n\n-----------------------------\n\n")

	scanner = bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
