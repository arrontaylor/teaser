package server

import "encoding/json"
import "github.com/arrontaylor/teaser/teaser"
import "net/http"
import "strconv"

var ProfileHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
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

	resp := make(map[string]interface{})

	resp["username"] = account.Username
	resp["sent"] = account.SentCount
	resp["friends"] = teaser.GetFriendsList(account.Username)
	resp["teases"] = teaser.GetTeasesList(account.Username)

	write, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"Something strange happened!\"}"))
		return
	}

	w.Write(write)
})
