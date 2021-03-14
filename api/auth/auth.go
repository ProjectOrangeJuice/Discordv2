package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/projectorangejuice/Discordv2/api"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", api.LoginHandler).Methods("POST")

	fmt.Println("Started auth server")

	log.Fatal(http.ListenAndServe(":8000", router))
}
