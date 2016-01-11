package server

import "github.com/arrontaylor/teaser/teaser"
import "net/http"
import "strconv"

var ReadTeaseHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	tousername := account.Username
	teaseid, _ := strconv.ParseInt(r.FormValue("teaseid"), 10, 0)

	teaser.ReadTease(int(teaseid), tousername)

	delete(teaser.TeasesLists, tousername)
	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
})
