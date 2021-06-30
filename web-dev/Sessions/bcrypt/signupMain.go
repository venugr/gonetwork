package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func getUser(r *http.Request) user {
	var u user

	cookie, err := r.Cookie("session")
	if err != nil {
		return u
	}

	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	return u
}

func isSingedUp(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil {
		return false
	}
	_, ok := dbSessions[cookie.Value]
	return ok
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/info", info)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	u := getUser(r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)

}

func info(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)

	if !isSingedUp(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "info.gohtml", u)

}

func signup(w http.ResponseWriter, r *http.Request) {

	if isSingedUp(r) {
		fmt.Println("User Signed already!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var u user

	if r.Method == http.MethodPost {

		fmt.Println("Post Method...processing.")
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")

		if _, ok := dbUsers[un]; ok {
			http.Error(w, "User already exists", http.StatusForbidden)
			return
		}

		uuid := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: uuid.String(),
		}

		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = un

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		u := user{un, bs, f, l}
		dbUsers[un] = u

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("Loading signup.gohtml....")
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}
