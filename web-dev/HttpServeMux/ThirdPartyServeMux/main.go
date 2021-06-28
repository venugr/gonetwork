package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()

	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyProcess)
	mux.GET("/user/:name", user)
	mux.GET("/blog/:category/:article", blogRead)

	http.ListenAndServe(":8080", mux)
}

func blogRead(respw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(respw, "READ CATEGORY: %s!\n", ps.ByName("category"))
	fmt.Fprintf(respw, "READ  ARTICLE: %s!\n", ps.ByName("article"))
}

func user(respw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(respw, "USER: %s!\n", ps.ByName("name"))
}

func index(respw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(respw, "index.gohtml", nil)
	HandleError(respw, err)
}

func about(respw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(respw, "about.gohtml", nil)
	HandleError(respw, err)
}

func contact(respw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(respw, "contact.gohtml", nil)
	HandleError(respw, err)
}

func apply(respw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(respw, "apply.gohtml", nil)
	HandleError(respw, err)
}

func applyProcess(respw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(respw, "applyProcess.gohtml", nil)
	HandleError(respw, err)
}

func HandleError(respw http.ResponseWriter, err error) {
	if err != nil {
		http.Error(respw, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
