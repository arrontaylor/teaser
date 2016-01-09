package server

import "net/http"
import "strconv"

var CreateTeaseHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	var account *Account

	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil {
			if username := Sessions[int(sessionId)]; username != nil {
				account = Accounts[*username]
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

	tease := CreateTease(fromusername, tousername)

	if tease == nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"Error creating tease\"}"))
	}

	delete(TeasesLists, tousername)
})
