package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type anytype int

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("serveHttp_data_host.gohtml"))
}

func (at anytype) ServeHTTP(respw http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	err = tpl.ExecuteTemplate(respw, "serveHttp_data_host.gohtml", data)
}

func main() {

	var at anytype
	http.ListenAndServe(":8080", at)

}
