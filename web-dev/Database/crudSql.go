package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {

	db, err = sql.Open("mysql", "admin:mypassword@tcp(mydbinstance.cotq313e2we9.us-east-1.rds.amazonaws.com:3306)/test02?charset=utf8")
	checkErr(err)

	defer db.Close()
	err = db.Ping()
	checkErr(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	checkErr(http.ListenAndServe(":8080", nil))

}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "At Index")

}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	// checkErr(err)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	// checkErr(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	n, err := res.RowsAffected()
	checkErr(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("Lella");`)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	checkErr(err)

	n, err := res.RowsAffected()
	checkErr(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	checkErr(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		checkErr(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Leela" WHERE name="Lella";`)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	checkErr(err)

	n, err := res.RowsAffected()
	checkErr(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Leela";`)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	checkErr(err)

	n, err := res.RowsAffected()
	checkErr(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	checkErr(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")

}

func amigos(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT aName FROM amigos")
	checkErr(err)

	defer rows.Close()

	var s, name string
	s = "Retrieved Records:\n"
	for rows.Next() {
		err := rows.Scan(&name)
		checkErr(err)
		s += name + "\n"
	}

	fmt.Fprintln(w, s)
}

func checkErr(err error) {

	if err != nil {
		log.Fatalln(err)
	}

}
