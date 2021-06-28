package main

import (
	"io"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", dognew)

	http.ListenAndServe(":8080", nil)

}

func dognew(respw http.ResponseWriter, req *http.Request) {

	respw.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(respw, `<img src="/doby.jpg">`)
}
