package main

import "github.com/arrontaylor/teaser/server"
import "github.com/arrontaylor/teaser/teaser"
import "github.com/arrontaylor/teaser/db"
import "net/http"
import "fmt"

func main() {
	links()
	routes()

	fmt.Println("Starting teaser server...")
	http.ListenAndServe(":8080", nil)
}

func links() {
	teaser.LoadAccount = db.LoadAccount
	teaser.CreateAccount = db.CreateAccount
	teaser.LoadFriendsList = db.LoadFriendsList
	teaser.AddFriendToList = db.AddFriendToList
	teaser.LoadTeasesList = db.LoadTeasesList
	teaser.CreateTease = db.CreateTease
	teaser.ReadTease = db.ReadTease
}

func routes() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.Handle("/login.go", server.LoginHandler)
	http.Handle("/logout.go", server.LogoutHandler)
	http.Handle("/signup.go", server.SignupHandler)
	http.Handle("/profile.json", server.ProfileHandler)
	http.Handle("/create-tease.go", server.CreateTeaseHandler)
	http.Handle("/read-tease.go", server.ReadTeaseHandler)
	http.Handle("/add-friend.go", server.AddFriendHandler)
}
