package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "admin:mypassword@tcp(mydbinstance.cotq313e2we9.us-east-1.rds.amazonaws.com:3306)/test02?charset=utf8")
	checkErr(err)

	defer db.Close()
	err = db.Ping()
	checkErr(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	checkErr(http.ListenAndServe(":8080", nil))

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Successfully completed.")
	checkErr(err)
}

func checkErr(err error) {

	if err != nil {
		fmt.Println(err)
	}

}
