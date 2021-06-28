package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", callName1)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func callName(respw http.ResponseWriter, req *http.Request) {
	key := req.FormValue("name")
	age := req.FormValue("age")
	key1 := req.FormValue("name")
	io.WriteString(respw, "Name: "+key+", Age: "+age+", "+key1)
}

func callName1(respw http.ResponseWriter, req *http.Request) {
	key := req.FormValue("name")

	respw.Header().Set("Content-Type", "text/html; charset=utf-8")

	// io.WriteString(respw, `
	// 	<form method="post">
	// 	   <input type="text" name="name">
	// 	   <input type="submit">
	// 	</form>
	// 	<br>`+key)

	io.WriteString(respw, `
		<form method="get">
		   <input type="text" name="name">
		   <input type="submit">
		</form>
		<br>`+key)
}
