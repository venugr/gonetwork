package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Template Name:", tpl.Name())
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println()
	for idx, lTpl := range tpl.Templates() {
		fmt.Printf("%d. %s\n", idx+1, lTpl.Name())
		fmt.Println("----------------")
		lTpl.Execute(os.Stdout, nil)
	}
	fmt.Println()

	fmt.Println("Template Name:", tpl.Name())
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
