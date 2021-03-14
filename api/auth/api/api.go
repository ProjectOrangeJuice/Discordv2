package api

import (
	"github.com/projectorangejuice/Discordv2/db"
	"net/http"
	"fmt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
a,b := db.DoesUserExist("test")
fmt.Printf("a :%v b:%v\n",a,b)
}