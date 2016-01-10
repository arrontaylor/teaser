package server

import "github.com/arrontaylor/teaser/teaser"
import "net/http"

var SignupHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	username := r.FormValue("username")
	password := HashPassword(r.FormValue("password"))
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	gender := r.FormValue("gender")

	account := teaser.CreateAccount(username, password, firstname, lastname, gender)
	if account == nil {
		w.WriteHeader(500)
		w.Write([]byte("Error creating account"))
		return
	} else {
		teaser.Accounts[account.Username] = account
	}

	LoginHandler(w, r)
})
