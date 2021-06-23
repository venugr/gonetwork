package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("data_map2.gohtml"))
}

func main() {

	sages := map[string]string{
		"1.India":    "Gandhi",
		"2.America":  "MLK",
		"3.Meditate": "Buddha",
		"4.Love":     "Jesus",
		"5.Prophet":  "Muhammad"}

	err := tpl.ExecuteTemplate(os.Stdout, "data_map2.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
