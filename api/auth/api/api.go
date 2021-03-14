package api

import (
	"io/ioutil"
	"net/http"

	"github.com/ProjectOrangeJuice/golib-logging/logging"
	"github.com/projectorangejuice/Discordv2/db"
)

// LoginHandler reads the body as the username and returns 200 OK if accepted, or 401
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logging.LogHTTP("loginHandler", "Failed to read body", err, w)
		return
	}

	ok, err := db.DoesUserExist(string(responseData))
	if err != nil {
		logging.LogHTTP("LoginHandler", "Failed to check username", err, w)
		return
	}

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
