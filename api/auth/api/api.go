package api

import (
	"net/http"

	"github.com/ProjectOrangeJuice/golib-logging/logging"
	"github.com/projectorangejuice/Discordv2/db"
)

// LoginHandler reads the body as the username and returns 200 OK if accepted, or 401
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Username empty", http.StatusInternalServerError)
	}

	ok, err := db.DoesUserExist(username)
	if err != nil {
		logging.LogHTTP("LoginHandler", "Failed to check username", err, w)
		return
	}

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
	}

	c := http.Cookie{Name: "orange-auth", Value: "Sir Dukkins"}
	http.SetCookie(w, &c)
}
