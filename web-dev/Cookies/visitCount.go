package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", readCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func readCookie(respw http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-visits")

	if err == http.ErrNoCookie {
		fmt.Println("Cookie not found!!")
		cookie = &http.Cookie{
			Name:  "my-visits",
			Value: "0",
			Path:  "/",
		}
	}

	i, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
		return
	}
	i++

	cookie.Value = strconv.Itoa(i)
	http.SetCookie(respw, cookie)

	io.WriteString(respw, cookie.Value)

}
