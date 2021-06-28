package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/set", setOnly)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/expire", setExpire)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(respw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(respw, `<h1><a href="/set">set a cookie</a></h1>`)
}

func setOnly(respw http.ResponseWriter, req *http.Request) {

	http.SetCookie(respw, &http.Cookie{
		Name:  "session",
		Value: "any-value",
		Path:  "/",
	})
	fmt.Fprintf(respw, `<h1><a href="/read">read</a></h1>`)
}

func readCookie(respw http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(respw, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(respw, `<h1>Your Cookie=%v</h1><br><h1><a href="/expire">expire</a></h1>`, cookie)
}

func setExpire(respw http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(respw, req, "/set", http.StatusSeeOther)
		return
	}

	cookie.MaxAge = -1
	http.SetCookie(respw, cookie)
	http.Redirect(respw, req, "/", http.StatusSeeOther)
}
