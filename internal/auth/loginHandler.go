package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func login(w http.ResponseWriter, r *http.Request) {
	key := []byte("super-secret-key")
	store := sessions.NewCookieStore(key)
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}
