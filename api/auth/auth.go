package main

import (
	"github.com/gorilla/mux"
	"github.com/projectorangejuice/Discordv2/api"
	"log"
	"net/http"
	"fmt"
)

func main(){

	router := mux.NewRouter()
	router.HandleFunc("/login", api.LoginHandler).Methods("GET")
	fmt.Println("running")
	log.Fatal(http.ListenAndServe(":8000",router))
}