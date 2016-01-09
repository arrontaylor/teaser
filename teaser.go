package main

import "github.com/arrontaylor/teaser/server"
import "github.com/arrontaylor/teaser/db"
import "net/http"
import "fmt"

func main() {
	server.LoadAccount = db.LoadAccount
	server.CreateAccount = db.CreateAccount
	server.LoadFriendsList = db.LoadFriendsList
	server.LoadTeasesList = db.LoadTeasesList
	server.CreateTease = db.CreateTease

	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.Handle("/login.go", server.LoginHandler)
	http.Handle("/signup.go", server.SignupHandler)
	http.Handle("/profile.json", server.ProfileHandler)
	http.Handle("/create-tease.go", server.CreateTeaseHandler)
	fmt.Println("Starting teaser server...")
	http.ListenAndServe(":8080", nil)
}
