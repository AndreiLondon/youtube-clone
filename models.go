package main

type User struct {
	Id        int
	Email     string
	Username  string
	Password  string
	Sessionid string
}

type SingCredentials struct {
	SignIn string
	SignUp string
}

type IndexObject struct {
	User *User
}
