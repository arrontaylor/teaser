package server

import "github.com/arrontaylor/teaser/teaser"
import "net/http"
import "strconv"

var AddFriendHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	username1 := account.Username
	username2 := r.FormValue("username2")

	friendaccount := teaser.GetAccount(username2)

	if friendaccount == nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"Cannot add friend '" + username2 + "': Account does not exist\"}"))
		return
	}

	teaser.AddFriendToList(username1, username2)

	delete(teaser.FriendsLists, username1)
})
