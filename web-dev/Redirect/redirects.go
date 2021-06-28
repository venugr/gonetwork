package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {

	http.HandleFunc("/", callRoot)

	http.HandleFunc("/one", callOneRedirect)
	http.HandleFunc("/onesubmit", callOneSubmit)

	http.HandleFunc("/two", callTwoRedirect)
	http.HandleFunc("/twosubmit", callTwoSubmit)

	http.HandleFunc("/moved", callMovedRedirect)
	http.HandleFunc("/movedsubmit", callMovedSubmit)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func callRoot(respw http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Request Method at Root:", req.Method)
}

func callOneRedirect(respw http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Request Method at One:", req.Method)
	http.Redirect(respw, req, "/", http.StatusSeeOther)
}

func callOneSubmit(respw http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Request Method at OneSubmit:", req.Method)
	tpl.ExecuteTemplate(respw, "redirectSeeOther.gohtml", nil)
}

func callTwoRedirect(respw http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Request Method at Two:", req.Method)
	http.Redirect(respw, req, "/", http.StatusTemporaryRedirect)
}

func callTwoSubmit(respw http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Request Method at TwoSubmit:", req.Method)
	tpl.ExecuteTemplate(respw, "redirectTemporary.gohtml", nil)
}

func callMovedRedirect(respw http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Request Method at Moved:", req.Method)
	http.Redirect(respw, req, "/", http.StatusMovedPermanently)
}

func callMovedSubmit(respw http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Request Method at MovedSubmit:", req.Method)
	tpl.ExecuteTemplate(respw, "redirectMovedPermanently.gohtml", nil)
}
