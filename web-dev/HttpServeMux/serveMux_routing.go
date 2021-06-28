package main

import (
	"io"
	"net/http"
)

type anytype int

func (at anytype) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(respw, "Dog Dog Dog...")
	case "/cat":
		io.WriteString(respw, "Cats are Cats..always")
	default:
		io.WriteString(respw, "Default Cats & Dogs....CDT")
	}
}

func main() {

	var at anytype
	http.ListenAndServe(":8080", at)
}
