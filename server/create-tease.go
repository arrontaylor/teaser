package server

import "github.com/arrontaylor/teaser/teaser"
import "net/http"
import "strconv"

var CreateTeaseHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	var account *teaser.Account

	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil {
			if username := Sessions[int(sessionId)]; username != nil {
				account = teaser.Accounts[*username]
			}
		}
	}

	if account == nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"Not logged in\"}"))
		return
	}

	fromusername := account.Username
	tousername := r.FormValue("tousername")

	tease := teaser.CreateTease(fromusername, tousername)

	if tease == nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"Error creating tease\"}"))
	}

	delete(teaser.TeasesLists, tousername)
	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
})
