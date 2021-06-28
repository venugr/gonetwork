package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tempgothml/*"))
}

func main() {

	http.HandleFunc("/", callProc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func callProc(respw http.ResponseWriter, req *http.Request) {

	s := ""

	fmt.Println(req.Method)

	if req.Method == http.MethodPost {
		mf, mh, err := req.FormFile("q")
		if err != nil {
			http.Error(respw, "Err1: "+err.Error(), http.StatusInternalServerError)
			fmt.Println(mf)
			return
		}

		mf.Close()

		bs, err := ioutil.ReadAll(mf)
		if err != nil {
			http.Error(respw, "Err2: "+err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		dst, err := os.Create(filepath.Join("./user", mh.Filename))
		if err != nil {
			http.Error(respw, "Err3: "+err.Error(), http.StatusInternalServerError)
			return
		}

		defer dst.Chdir()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(respw, "Err4: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	respw.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(respw, "index.gohtml", s)

}
