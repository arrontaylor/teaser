package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/arrontaylor/teaser/db"
	"github.com/arrontaylor/teaser/server"
	"github.com/arrontaylor/teaser/teaser"
	"net/http"
)

func main() {
	links()
	routes()

	log.SetLevel(log.DebugLevel)

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
	teaser.LoadFrequentsList = db.LoadFrequentsList
	teaser.LoadRank = db.LoadRank
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
