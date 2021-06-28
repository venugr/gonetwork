package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", setCookie)
	http.HandleFunc("/set", setOnly)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/abundance", setMultiCookies)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func setCookie(respw http.ResponseWriter, req *http.Request) {

	http.SetCookie(respw, &http.Cookie{
		Name:  "my-cookie",
		Value: "my-first-cookie-in-go",
		Path:  "/",
	})
	fmt.Fprintln(respw, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(respw, "in chrome go to: dev tools / application / cookies")
}

func setOnly(respw http.ResponseWriter, req *http.Request) {

	http.SetCookie(respw, &http.Cookie{
		Name:  "setOnly",
		Value: "1234567",
		Path:  "/readonly",
	})
	fmt.Fprintln(respw, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(respw, "in chrome go to: dev tools / application / cookies")
}

func readCookie1(respw http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-cookie")
	if err != nil {
		fmt.Fprintln(respw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(respw, "Your Cookie #2: ", cookie)
}

func readCookie(respw http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(respw, "Your Cookie #1: ", cookie)
	}

	cookie, err = req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(respw, "Your Cookie #2: ", cookie)
	}

	cookie, err = req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(respw, "Your Cookie #3: ", cookie)
	}

}

func setMultiCookies(respw http.ResponseWriter, req *http.Request) {

	http.SetCookie(respw, &http.Cookie{
		Name:  "general",
		Value: "general-cookie-in-go",
	})

	http.SetCookie(respw, &http.Cookie{
		Name:  "specific",
		Value: "specific-cookie-in-go",
	})

	fmt.Fprintln(respw, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(respw, "in chrome go to: dev tools / application / cookies")
}
