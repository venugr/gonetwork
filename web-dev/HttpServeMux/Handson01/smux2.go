package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("smux2.gohtml"))
}

func defIndex(respw http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(respw, "smux2.gohtml", "Default")
}

func callDog(respw http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(respw, "Dog Dogs Dogger....")
	//io.WriteString(respw, "Dog Dogs Dogger....")
	tpl.ExecuteTemplate(respw, "smux2.gohtml", "Dog Dogs Dogger....")
}

func callMe(respw http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(respw, "Thanks for calling - VenuLella")
	tpl.ExecuteTemplate(respw, "smux2.gohtml", "Thanks for calling - VenuLella")
}

func main1() {

	http.HandleFunc("/", defIndex)
	http.HandleFunc("/index", defIndex)
	http.HandleFunc("/index.html", defIndex)
	http.HandleFunc("/dog/", callDog)
	http.HandleFunc("/me", callMe)

	http.ListenAndServe(":8080", nil)
}

func main() {

	http.Handle("/", http.HandlerFunc(defIndex))
	http.Handle("/index", http.HandlerFunc(defIndex))
	http.Handle("/index.html", http.HandlerFunc(defIndex))
	http.Handle("/dog/", http.HandlerFunc(callDog))
	http.Handle("/me", http.HandlerFunc(callMe))

	http.ListenAndServe(":8080", nil)
}
