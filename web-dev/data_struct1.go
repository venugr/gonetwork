package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("data_struct1.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no believes"}

	err := tpl.ExecuteTemplate(os.Stdout, "data_struct1.gohtml", buddha)
	if err != nil {
		log.Fatalln(err)
	}

}
