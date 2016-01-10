package server

import "github.com/arrontaylor/teaser/teaser"
import "net/http"
import "strconv"

var LogoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// redirect if already logged in

	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil {
			if username := Sessions[int(sessionId)]; username != nil {
				delete(teaser.Accounts, *username)
			}

			delete(Sessions, int(sessionId))
		}
	}

	w.Header().Set("Set-Cookie", "SessionId=0")
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
})
