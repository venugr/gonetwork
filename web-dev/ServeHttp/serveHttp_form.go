package main

import (
	"html/template"
	"log"
	"net/http"
)

type anytype int

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("serveHttp_form.gohtml"))
}

func (at anytype) ServeHTTP(respw http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(respw, "serveHttp_form.gohtml", req.Form)
}

func main() {

	var at anytype
	http.ListenAndServe(":8080", at)

}
