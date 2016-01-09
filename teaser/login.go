package teaser

import "net/http"
import "strconv"

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
	}

	// redirect if already logged in
	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil && Sessions[int(sessionId)] != nil {
			http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
			return
		}
	}

	username := r.FormValue("username")
	password := HashPassword(r.FormValue("password"))

	account := GetAccount(username)

	if account == nil {
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
