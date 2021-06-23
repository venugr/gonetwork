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

	tpl = template.Must(template.ParseFiles("serveHttp_data_url.gohtml"))
}

func (at anytype) ServeHTTP(respw http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
	}{
		req.Method,
		req.URL,
		req.Form,
	}

	err = tpl.ExecuteTemplate(respw, "serveHttp_data_url.gohtml", data)
}

func main() {

	var at anytype
	http.ListenAndServe(":8080", at)

}
