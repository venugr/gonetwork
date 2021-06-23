package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("data_nested*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "data_nested1.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

}
