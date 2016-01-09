package teaser

import "net/http"

var SignupHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	username := r.FormValue("username")
	password := HashPassword(r.FormValue("password"))

	account := CreateAccount(username, password)
	if account == nil {
		w.WriteHeader(500)
		w.Write([]byte("Error creating account"))
		return
	} else {
		Accounts[account.Username] = account
	}

	LoginHandler(w, r)
})
