package server

import "net/http"
import "strconv"
import "fmt"
import "encoding/json"

var ProfileHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
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

	resp := make(map[string]interface{})

	resp["username"] = account.Username
	resp["friends"] = GetFriendsList(account.Username)
	resp["teases"] = GetTeasesList(account.Username)

	fmt.Println(resp["teases"])

	write, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"Something strange happened!\"}"))
		return
	}

	w.Write(write)
})
