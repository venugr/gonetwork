package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("data_and_logic.gohtml"))
}

func main() {

	u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	v1 := user{}
	users := []user{u1, v1, u2, u3}

	_ = users

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
