package main

import (
	"fmt"
	"os"
	"log"
	"database/sql"
	"html/template"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Pwn struct {
	ID          int
	Email       string
	Password    string
}

var tpl *template.Template

var db *sql.DB

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	var err error

	db, err = sql.Open("mysql", "root:password.@tcp(localhost:3306)/data")
	err = db.Ping()
  	errCheck(err)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	http.HandleFunc("/", pwnSearchHandler)

	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println(os.Getwd())
}

func errCheck(err error){
  if err != nil{
    log.Fatal(err)
  }
}

func pwnSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "index.html", nil)
		return
	}
	r.ParseForm()
	name := r.FormValue("pwnEmail")

	stmt := "SELECT * FROM pwned WHERE email = ?;"

	rows, err := db.Query(stmt, name)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var pwned []Pwn

	for rows.Next() {
		var p Pwn

		err = rows.Scan(&p.ID, &p.Email, &p.Password)
		if err != nil {
			panic(err)
		}

		pwned = append(pwned, p)
	}
	tpl.ExecuteTemplate(w, "index.html", pwned)
}
