package server

import (
	log "github.com/Sirupsen/logrus"
	"github.com/arrontaylor/teaser/teaser"
	"net/http"
	"strconv"
)

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Debug("Login request not \"POST\": ", r)
		w.WriteHeader(404)
	}

	// redirect if already logged in
	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil && Sessions[int(sessionId)] != nil {
			log.Debug("Login redirect, session exists: ", Sessions[int(sessionId)])
			http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
			return
		}
	}

	username := r.FormValue("username")
	password := HashPassword(r.FormValue("password"))

	account := teaser.GetAccount(username)

	if account == nil {
		log.Debug("Login failed. Account does not exist: " + username)
		http.Redirect(w, r, "/signup", http.StatusMovedPermanently)
		return
	}

	if account.Password != password {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		return
	}

	if account.SessionId == 0 {
		account.SessionId = CreateSessionId()
	}

	Sessions[account.SessionId] = &account.Username

	w.Header().Set("Set-Cookie", "SessionId="+strconv.Itoa(account.SessionId))
	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
})
