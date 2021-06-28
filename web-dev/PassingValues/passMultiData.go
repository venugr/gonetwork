package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", callProc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func callProc(respw http.ResponseWriter, req *http.Request) {

	fName := req.FormValue("first")
	lName := req.FormValue("last")
	subOk := req.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(respw, "index.gohtml", person{fName, lName, subOk})
	if err != nil {
		http.Error(respw, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}

}
