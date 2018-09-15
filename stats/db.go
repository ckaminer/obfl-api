package stats

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "obfl"
)

var OBFLDB DB

type Owner struct {
	ID    int      `json:"-"`
	Name  string   `json:"name"`
	Teams []string `json:"teams"`
}

type DB struct {
	*sql.DB
}

func GetAllOwners() ([]Owner, error) {
	query := `SELECT * FROM owners;`

	rows, err := OBFLDB.Query(query)
	if err != nil {
		log.Println("Error running query: ", err.Error())
		return nil, err
	}

	var owners []Owner
	for rows.Next() {
		var owner Owner
		err := rows.Scan(&owner.ID, &owner.Name)
		if err != nil {
			log.Println("Error scanning owner info: ", err.Error())
			return nil, err
		}

		owners = append(owners, owner)
	}

	return owners, nil
}
