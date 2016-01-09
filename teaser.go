package main

import "github.com/arrontaylor/teaser/teaser"
import "net/http"
import "fmt"

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/login.go", teaser.LoginHandler)
	http.Handle("/signup.go", teaser.SignupHandler)
	fmt.Println("Starting teaser server...")
	http.ListenAndServe(":8080", nil)
}
