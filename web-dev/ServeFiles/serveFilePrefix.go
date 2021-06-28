package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/dog/", dognew)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8080", nil)

}

func dognew(respw http.ResponseWriter, req *http.Request) {
	respw.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(respw, `<img src="/resources/doby.jpg">`)
}
