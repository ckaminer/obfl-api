package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ckaminer/obfl-api/server"
	"github.com/ckaminer/obfl-api/stats"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		stats.Host, stats.Port, stats.User, stats.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Error opening DB Connection: ", err.Error())
	}
	defer db.Close()

	stats.OBFLDB = stats.DB{db}

	server.StartServer()
}
