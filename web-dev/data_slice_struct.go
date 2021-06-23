package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("data_slice_struct.gohtml"))
}

func main() {

	sg1 := sage{
		Name:  "Buddha",
		Motto: "The belief of no believes",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	sages := []sage{sg1, gandhi, mlk, jesus, muhammad}

	err := tpl.ExecuteTemplate(os.Stdout, "data_slice_struct.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}