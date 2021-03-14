package db

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
)

var con *sql.DB

func init(){
	// Create the database if required
	var err error
	con, err = sql.Open("sqlite3", "db.sqlite")

	if err != nil { 
		log.Fatalf("Failed to open sqlite")
	}

	statement, _ := con.Prepare(`CREATE TABLE IF NOT EXISTS users 
			(id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT UNIQUE)`)
	statement.Exec()
}


func DoesUserExist(username string) (bool, error){
	stmt, _ := con.Prepare(`SELECT id
			FROM users WHERE username = 'test' LIMIT 1;`)
			rows, err := stmt.Query()
			if err != nil {
				return false, err
			}


			if rows.Next(){
				return true,nil
			}

			return false,nil
}
