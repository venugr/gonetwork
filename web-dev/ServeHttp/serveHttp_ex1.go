package main

import (
	"fmt"
	"net/http"
)

type anytype int

func (at anytype) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(respw, "Any code you want to write - VenuLella")
}

func main() {

	var at anytype

	http.ListenAndServe(":8080", at)
}
