package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdateMDY":      monthDayYear,
	"fdateDMY":      dayMonthYear,
	"fdateDDMMYYYY": ddMmmYyyy,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

func ddMmmYyyy(t time.Time) string {
	return t.Format("02-Jan-2006")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("data_func2.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "data_func2.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
