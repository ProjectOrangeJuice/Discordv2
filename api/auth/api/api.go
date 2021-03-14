package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/projectorangejuice/Discordv2/db"
)

// LoginHandler reads the body as the username and returns 200 OK if accepted, or 401
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read body for username in login handler, %s", err)
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	ok, err := db.DoesUserExist(string(responseData))
	if err != nil {
		log.Printf("Failed to check username, %s", err)
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
