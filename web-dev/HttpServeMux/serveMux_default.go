package main

import (
	"fmt"
	"io"
	"net/http"
)

type anycat int
type anydog int
type anyany int

func (ac anycat) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
	io.WriteString(respw, fmt.Sprintf("Cats Cats..and CCCCCCCCCATS.....%s", req.URL.Path))
}

func (ac anydog) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
	io.WriteString(respw, fmt.Sprintf("DOGS DOGS....DOOOOOOOOOOOGS: %s", req.URL.Path))
}

func (ac anyany) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
	io.WriteString(respw, fmt.Sprintf("ANY ANY ANY....ANYYYYYYYYYY: %s", req.URL.Path))
}

func acf(respw http.ResponseWriter, req *http.Request) {
	io.WriteString(respw, fmt.Sprintf("ACF: Cats Cats..and CCCCCCCCCATS.....%s", req.URL.Path))
}

func adf(respw http.ResponseWriter, req *http.Request) {
	io.WriteString(respw, fmt.Sprintf("ADF: DOGS DOGS....DOOOOOOOOOOOGS: %s", req.URL.Path))
}

func aaf(respw http.ResponseWriter, req *http.Request) {
	io.WriteString(respw, fmt.Sprintf("AAF: --- ANY ANY ANY....ANYYYYYYYYYY: %s", req.URL.Path))
}

func main1() {

	var ac anycat
	var ad anydog
	var aa anyany

	mux := http.NewServeMux()
	mux.Handle("/dog/", ad)
	mux.Handle("/cat", ac)
	mux.Handle("/", aa)

	http.ListenAndServe(":8080", mux)
}

func main2() {

	var ac anycat
	var ad anydog
	var aa anyany

	//mux := http.NewServeMux()
	http.Handle("/dog/", ad)
	http.Handle("/cat", ac)
	http.Handle("/", aa)

	http.ListenAndServe(":8080", nil)
}

func main3() {
	http.HandleFunc("/dog/", adf)
	http.HandleFunc("/cat", acf)
	http.HandleFunc("/", aaf)

	http.ListenAndServe(":8080", nil)
}

func main() {
	http.Handle("/dog/", http.HandlerFunc(adf))
	http.Handle("/cat", http.HandlerFunc(acf))
	http.Handle("/", http.HandlerFunc(aaf))

	http.ListenAndServe(":8080", nil)
}
