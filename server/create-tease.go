package server

import "github.com/arrontaylor/teaser/teaser"
import "net/http"
import "strconv"

var CreateTeaseHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	var fromaccount *teaser.Account

	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil {
			if username := Sessions[int(sessionId)]; username != nil {
				fromaccount = teaser.Accounts[*username]
			}
		}
	}

	if fromaccount == nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"Not logged in\"}"))
		return
	}

	fromusername := fromaccount.Username
	tousername := r.FormValue("tousername")

	toaccount := teaser.GetAccount(tousername)

	if toaccount == nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"Cannot send tease to '" + tousername + "': Account does not exist\"}"))
		return
	}

	tease := teaser.CreateTease(fromusername, tousername)

	if tease == nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\":\"Error creating tease\"}"))
		return
	}

	fromaccount.SentCount++

	delete(teaser.TeasesLists, tousername)
	delete(teaser.FrequentsLists, fromusername)
	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
})
