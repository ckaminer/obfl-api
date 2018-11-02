package stats

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	Host   = "localhost"
	Port   = 5432
	User   = "charleskaminer"
	DbName = "obfl"
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
		log.Println("Error finding owners: ", err.Error())
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

func CreateTeam(teamName string) (int, error) {
	query := `INSERT INTO teams (name) values ($1) RETURNING team_id;`

	var teamID int
	err := OBFLDB.QueryRow(query, teamName).Scan(&teamID)
	if err != nil {
		log.Println("Error creating team: ", err.Error())
		return 0, err
	}

	return teamID, nil
}

func FindTeam(teamName string) (int, error) {
	query := `SELECT team_id FROM teams WHERE name = $1;`

	var id int
	err := OBFLDB.QueryRow(query, teamName).Scan(&id)
	if err != nil {
		log.Println("Error finding team: ", err.Error())
		return 0, err
	}

	return id, nil
}

func FindOrCreateTeam(teamName string) (int, error) {
	id, err := FindTeam(teamName)
	if err != nil {
		if err == sql.ErrNoRows {
			return CreateTeam(teamName)
		}
		return 0, err
	}

	return id, err
}

func CreateOwner(ownerName string) (int, error) {
	query := `INSERT INTO owners (name) values ($1) RETURNING owner_id;`

	var ownerID int
	err := OBFLDB.QueryRow(query, ownerName).Scan(&ownerID)
	if err != nil {
		log.Println("Error creating owner: ", err.Error())
		return 0, err
	}

	return ownerID, nil
}

func FindOwner(ownerName string) (int, error) {
	query := `SELECT owner_id FROM owners WHERE name = $1;`

	var id int
	err := OBFLDB.QueryRow(query, ownerName).Scan(&id)
	if err != nil {
		log.Println("Error finding owner: ", err.Error())
		return 0, err
	}

	return id, nil
}

func FindOrCreateOwner(ownerName string) (int, error) {
	id, err := FindOwner(ownerName)
	if err != nil {
		if err == sql.ErrNoRows {
			return CreateOwner(ownerName)
		}
		return 0, err
	}

	return id, err
}

func CreateTeamOwner(teamID, ownerID int) (int, error) {
	query := `INSERT INTO teams_to_owners (team_id, owner_id) values ($1, $2) RETURNING team_owner_id;`

	var toID int
	err := OBFLDB.QueryRow(query, teamID, ownerID).Scan(&toID)
	if err != nil {
		log.Println("Error creating team owner: ", err.Error())
		return 0, err
	}

	return toID, nil
}

func FindTeamOwner(teamID, ownerID int) (int, error) {
	query := `SELECT team_owner_id FROM teams_to_owners WHERE team_id = $1 AND owner_id = $2;`

	var id int
	err := OBFLDB.QueryRow(query, teamID, ownerID).Scan(&id)
	if err != nil {
		log.Println("Error finding team owner: ", err.Error())
		return 0, err
	}

	return id, nil
}

func FindOrCreateTeamOwner(teamID, ownerID int) (int, error) {
	id, err := FindTeamOwner(teamID, ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return CreateTeamOwner(teamID, ownerID)
		}
		return 0, err
	}

	return id, err
}
