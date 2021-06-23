package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("data_global_func1.gohtml"))
}

func main() {

	xs := []string{"zero", "one", "two", "three", "four", "five"}

	err := tpl.ExecuteTemplate(os.Stdout, "data_global_func1.gohtml", xs)
	if err != nil {
		log.Fatalln(err)
	}
}
