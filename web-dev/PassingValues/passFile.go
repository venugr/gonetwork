package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", callProc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func callProc(respw http.ResponseWriter, req *http.Request) {

	s := ""

	fmt.Println(req.Method)

	if req.Method == http.MethodPost {
		mf, _, err := req.FormFile("q")
		if err != nil {
			http.Error(respw, "Err1: "+err.Error(), http.StatusInternalServerError)
			fmt.Println(mf)
			return
		}

		mf.Close()
		//fmt.Println("\nfile:", mf, "\nheader:", mh, "\nerr", err)
		bs, err := ioutil.ReadAll(mf)
		if err != nil {
			http.Error(respw, "Err2: "+err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	respw.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(respw, `<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br><plaintext>`+s+"</plaintext>")

}
