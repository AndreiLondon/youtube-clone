package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var invalidCredentialsFlagSignUp = ""
var invalidCredentialsFlagSignIn = ""
var SESSION_ID = "SESSION_ID"

func main() {
	dbLocal, err := sql.Open("sqlite3", "./youtube.db")
	if err != nil {
		log.Fatal(err)
	}
	db = dbLocal
	defer db.Close()
	createTables()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/sign", signHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signout", signoutHandler)
	fmt.Println("Server start at port :9000")
	http.ListenAndServe(":9000", nil)
}
func createTables() {
	err := crerateUsersTable()
	if err != nil {
		log.Fatal(err)
	}
}
func showError(w http.ResponseWriter, code int, message string) {
	templ, err := template.ParseFiles("templates/error.html")
	w.WriteHeader(code)
	if err != nil {
		fmt.Fprint(w, "500 Internal Server Error")
		return
	}
	templ.Execute(w, message)
}
