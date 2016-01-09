package main

import "github.com/arrontaylor/teaser/teaser"
import "github.com/arrontaylor/teaser/db"
import "net/http"
import "fmt"

func main() {
	teaser.LoadAccount = db.LoadAccount

	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.Handle("/login.go", teaser.LoginHandler)
	http.Handle("/signup.go", teaser.SignupHandler)
	fmt.Println("Starting teaser server...")
	http.ListenAndServe(":8080", nil)
}
