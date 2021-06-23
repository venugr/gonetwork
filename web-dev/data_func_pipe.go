package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"os"
	"time"
)

var tpl *template.Template

func double(x int) int {
	return x + x
}

func square(f int) float64 {
	return math.Pow(float64(f), 2)
}

func squareRoot(f float64) float64 {
	return math.Sqrt(f)
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

func numToText(n int) string {

	txt := ""
	if n <= 10 {
		txt = "Num is greater than 10"
	}

	if n > 10 && n <= 100 {
		txt = "Num is greater than 10 and less than or equal to 100"
	}

	if n > 100 {
		txt = "Num is greater than 100"
	}

	return txt
}

var fm = template.FuncMap{
	"fdbl":          double,
	"fsqr":          square,
	"fsqrt":         squareRoot,
	"fn2t":          numToText,
	"fdateMDY":      monthDayYear,
	"fdateDMY":      dayMonthYear,
	"fdateDDMMYYYY": ddMmmYyyy,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("data_func_pipe.gohtml", "data_func_pipe2.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "data_func_pipe.gohtml", 112)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n------------------------\n")

	err = tpl.ExecuteTemplate(os.Stdout, "data_func_pipe2.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}
