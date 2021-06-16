package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("tpl.gohtml", "tpl1.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("----------WITH DATA----------")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("----------WITH VARIABLE VALUE INT----------")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", 1142)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("----------WITH VARIABLE VALUE STRING----------")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", "'THIS IS A TEXT'")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("----------WITHOUT VARIABLE VALUE----------")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
