package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName  string
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", getBar)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		uuid := uuid.NewV4()

		cookie = &http.Cookie{
			Name:  "session",
			Value: uuid.String(),
		}

		fmt.Println(cookie)
		http.SetCookie(w, cookie)

	}

	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	if r.Method == http.MethodPost {
		u.UserName = r.FormValue("username")
		u.FirstName = r.FormValue("firstname")
		u.LastName = r.FormValue("lastname")

		dbSessions[cookie.Value] = u.UserName
		dbUsers[u.UserName] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func getBar(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "getbar.gohtml", u)

}
