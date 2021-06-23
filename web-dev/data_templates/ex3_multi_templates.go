package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

const (
	ClrBlack = "\x1b[30;1m"

	ClrRed = "\x1b[31;1m"

	ClrGreen = "\x1b[32;1m"

	ClrYellow = "\x1b[33;1m"

	ClrBlue = "\x1b[34;1m"

	ClrMagenta = "\x1b[35;1m"

	ClrCyan = "\x1b[36;1m"

	ClrWhite = "\x1b[37;1m"

	ClrUnColor = "\x1b[0m"
)

func main() {

	colors := []string{
		"\x1b[31;1m",
		"\x1b[33;1m",
		"\x1b[34;1m",
		"\x1b[35;1m",
		"\x1b[36;1m",
		"\x1b[37;1m",
		"\x1b[32;1m",
	}

	tpl, err := template.ParseFiles("one.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("two.gohtml", "three.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	lTpl := tpl.Lookup("three.gohtml")
	lTpl.Execute(os.Stdout, nil)

	fmt.Println()
	for idx, tpl := range tpl.Templates() {
		// tColor := ClrRed
		// if idx%2 == 0 {
		// 	tColor = ClrBlue
		// }
		fmt.Printf("Templ Name:%s %s %s\n", colors[idx], tpl.Name(), ClrUnColor)
		err = tpl.ExecuteTemplate(os.Stdout, tpl.Name(), nil)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}
	fmt.Println()

}
