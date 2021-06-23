package main

import (
	"fmt"
	"net/http"
)

type anytype int

func (m anytype) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("VenuLell-Key", "this is VenuLellaG")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func - Venu</h1>")
}

func main() {
	var d anytype
	http.ListenAndServe(":8080", d)
}
