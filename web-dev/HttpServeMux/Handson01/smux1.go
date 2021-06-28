package main

import (
	"fmt"
	"io"
	"net/http"
)

func defIndex(respw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(respw, "Default index....")
}

func callDog(respw http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(respw, "Dog Dogs Dogger....")
	io.WriteString(respw, "Dog Dogs Dogger....")
}

func callMe(respw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(respw, "Thanks for calling - VenuLella")
}

func main() {

	http.HandleFunc("/", defIndex)
	http.HandleFunc("/index", defIndex)
	http.HandleFunc("/index.html", defIndex)
	http.HandleFunc("/dog/", callDog)
	http.HandleFunc("/me", callMe)

	http.ListenAndServe(":8080", nil)
}
