package teaser

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/login.go", LoginHandler)
	http.Handle("/signup.go", SignupHandler)
	http.ListenAndServe(":80", nil)
}
