package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	Role     string
}

var dbRegSessions = map[string]string{}
var dbLoginSessions = map[string]string{}
var dbUser = map[string]user{}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", getInfo)
	http.HandleFunc("/register", doRegister)
	http.HandleFunc("/login", doLogin)
	http.HandleFunc("/logout", doLogout)
	http.HandleFunc("/play", doPlay)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func getInfo(w http.ResponseWriter, r *http.Request) {

	rUid, lUid, rOk, lOk := getRegLoginDetails(r)
	infoData := struct {
		RegUid   string
		RegOk    bool
		LoginUid string
		LoginOk  bool
		Role     string
	}{
		rUid, rOk, lUid, lOk, dbUser[lUid].Role,
	}

	ckVal, _, ok := getLoginCookie(r)

	if ok {
		cookie := &http.Cookie{
			Name:   "loginid",
			Value:  ckVal,
			MaxAge: 60 * 3,
		}
		http.SetCookie(w, cookie)
	}

	// fmt.Println(infoData)
	tpl.ExecuteTemplate(w, "info.gohtml", infoData)

}

func doPlay(w http.ResponseWriter, r *http.Request) {

	rUid, lUid, rOk, lOk := getRegLoginDetails(r)
	infoData := struct {
		RegUid   string
		RegOk    bool
		LoginUid string
		LoginOk  bool
		Role     string
	}{
		rUid, rOk, lUid, lOk, dbUser[lUid].Role,
	}

	// fmt.Println(infoData)
	if !lOk {
		http.Error(w, "Not authorised to access this page...Register/Login", http.StatusForbidden)
		return
	}

	if dbUser[lUid].Role != "007" {
		http.Error(w, fmt.Sprintf("Insufficient role privlleges for '%v'", infoData.Role), http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(w, "play.gohtml", infoData)
}

func doLogout(w http.ResponseWriter, r *http.Request) {

	cookieValue, userName, ok := getLoginCookie(r)
	if !ok {
		http.Error(w, fmt.Sprintf("Something went wrong while logout for user: '%v", userName), http.StatusForbidden)
		return
	}

	cookie := &http.Cookie{
		Name:   "loginid",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
	delete(dbLoginSessions, cookieValue)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func doLogin(w http.ResponseWriter, r *http.Request) {
	lUid, lOk := getLoginDetails(r)
	infoData := struct {
		RegUid string
		RegOk  bool
	}{
		lUid, lOk,
	}

	if r.Method == http.MethodPost {
		userName := r.FormValue("username")
		password := r.FormValue("password")

		if !userExists(userName) {
			http.Error(w, fmt.Sprintf("User '%v' is not registered, Pl register!", userName), http.StatusForbidden)
			return
		}

		u := dbUser[userName]
		if !passwordMatched(u, password) {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		uuid := uuid.NewV4()
		cookie := &http.Cookie{
			Name:   "loginid",
			Value:  uuid.String(),
			MaxAge: 60 * 3,
		}

		http.SetCookie(w, cookie)
		dbLoginSessions[cookie.Value] = userName

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", infoData)
}

func doRegister(w http.ResponseWriter, r *http.Request) {
	rUid, rOk := getRegDetails(r)
	infoData := struct {
		RegUid string
		RegOk  bool
	}{
		rUid, rOk,
	}

	if r.Method == http.MethodPost {
		userName := r.FormValue("username")
		password := r.FormValue("password")
		role := r.FormValue("role")

		if userExists(userName) {
			http.Error(w, fmt.Sprintf("User '%v' already exists!", userName), http.StatusForbidden)
			return
		}

		bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		u := user{userName, bs, role}
		dbRegSessions[userName] = userName
		dbUser[userName] = u

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "register.gohtml", infoData)
}

func userExists(userName string) bool {
	_, ok := dbUser[userName]
	return ok
}

func passwordMatched(u user, password string) bool {

	err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		// http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		// return

		return false
	}

	return true
}

func getRegLoginDetails(r *http.Request) (string, string, bool, bool) {

	rUid, rOk := getRegDetails(r)
	lUid, lOk := getLoginDetails(r)

	return rUid, lUid, rOk, lOk

}

func getRegDetails(r *http.Request) (string, bool) {

	cookie, err := r.Cookie("reg-id")
	if err != nil {
		return "", false
	}

	if regId, ok := dbRegSessions[cookie.Value]; ok {
		return regId, ok
	}

	return "", false
}

func getLoginDetails(r *http.Request) (string, bool) {

	cookie, err := r.Cookie("loginid")
	if err != nil {
		return "", false
	}

	if regId, ok := dbLoginSessions[cookie.Value]; ok {
		return regId, ok
	}

	return "", false
}

func getLoginCookie(r *http.Request) (string, string, bool) {

	cookie, err := r.Cookie("loginid")
	if err != nil {
		return "", "", false
	}

	if un, ok := dbLoginSessions[cookie.Value]; ok {
		return cookie.Value, un, ok
	}

	return "", "", false
}
