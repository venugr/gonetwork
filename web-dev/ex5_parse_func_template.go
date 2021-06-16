package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	for idx, lTpl := range tpl.Templates() {
		fmt.Printf("%d. %s\n", idx+1, lTpl.Name())
	}

	err := tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println()
	for idx, lTpl := range tpl.Templates() {
		fmt.Printf("%d. %s\n", idx+1, lTpl.Name())
		tpl.ExecuteTemplate(os.Stdout, lTpl.Name(), nil)
		fmt.Println()
	}
	fmt.Println()
}
