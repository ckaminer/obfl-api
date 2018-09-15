package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ckaminer/obfl/server"
	"github.com/ckaminer/obfl/stats"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "obfl"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Error opening DB Connection: ", err.Error())
	}
	defer db.Close()

	stats.OBFLDB = stats.DB{db}

	server.StartServer()
}
