package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/doby", dogDoby)
	http.HandleFunc("/doby.jpg", dogServe)
	http.HandleFunc("/pic", dogPic)
	http.ListenAndServe(":8080", nil)

}

func dog(respw http.ResponseWriter, req *http.Request) {
	respw.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(respw, `
	<h1>From Wiki</h1>
	<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}

func dogDoby(respw http.ResponseWriter, req *http.Request) {
	respw.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(respw, `
	<h1>From File</h1>
	<!--not serving from our server-->
	<img src="/doby.jpg">
	`)
}

func dogPic(respw http.ResponseWriter, req *http.Request) {
	f, err := os.Open("doby.jpg")
	if err != nil {
		http.Error(respw, "File not found", 404)
		return
	}
	defer f.Close()

	io.Copy(respw, f)
}

func dogServe(respw http.ResponseWriter, req *http.Request) {
	f, err := os.Open("doby.jpg")
	if err != nil {
		http.Error(respw, "File not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(respw, "File not found", 404)
		return
	}

	http.ServeContent(respw, req, f.Name(), fi.ModTime(), f)
}
