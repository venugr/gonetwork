package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("data_map1.gohtml"))
}

func main() {

	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "data_map1.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
