package main

import (
	"html/template"
	"net/http"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		showError(w, 404, "404 Page not Found")
		return
	}
	invalidCredentialsFlagSignUp = ""
	invalidCredentialsFlagSignIn = ""
	indexObject := IndexObject{}
	// sessionId := getCookie(r)
	// videos := getUserBySessionId(sessionId)

	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	err = templ.Execute(w, indexObject)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	printUsers()
}
func signHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/sign.html")
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	singCredentials := SingCredentials{}
	singCredentials.SignIn = invalidCredentialsFlagSignIn
	singCredentials.SignUp = invalidCredentialsFlagSignUp

	err = templ.Execute(w, singCredentials)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
}
func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	username := strings.TrimSpace(r.FormValue("signup_username"))
	email := strings.TrimSpace(r.FormValue("signup_email"))
	password := strings.TrimSpace(r.FormValue("signup_password"))
	if username == "" || email == "" || password == "" || len(username) > 40 || len(password) < 3 {
		invalidCredentialsFlagSignUp = "Invalid username or password"
		invalidCredentialsFlagSignIn = ""
		http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
		return
	}
	sessionId := generateSessionId()
	err := saveUser(username, email, encrypt(password), sessionId)
	if err != nil {
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed:") {
			invalidCredentialsFlagSignUp = "User name or email alreay in use"
			invalidCredentialsFlagSignIn = ""
			http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
			return
		}
		showError(w, 500, "500 Internal Server Error. Error while working with database")
		return
	}
	setCookie(w, sessionId)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	email := r.FormValue("login_email")
	password := r.FormValue("login_password")
	user, err := checkUser(email, password)
	if err != nil {
		showError(w, 500, "500 Internal Server Error. Error while working with database")
		return
	}
	if err == nil && user == nil {
		invalidCredentialsFlagSignIn = "Invalid email or password"
		invalidCredentialsFlagSignUp = ""
		http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
		return
	}
	if user != nil {
		sessionId := generateSessionId()
		setCookie(w, sessionId)
		err := setSessionId(user, sessionId)
		if err != nil {
			showError(w, 500, "500 Internal Server Error")
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func signoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getCookie(r)
	if sessionId != "" {
		err := resetSessionId(sessionId)
		if err != nil {
			showError(w, 500, "500 Internal Server Error. Error while working with database")
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
