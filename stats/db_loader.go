package stats

import (
	"log"
)

func LoadTeamOwners(filePath string) error {
	teamMap, err := ReadTeamOwners(filePath)
	if err != nil {
		log.Println("Error Loading Data:: ", err.Error())
		return err
	} else {
		for owner, team := range teamMap {
			oID, err := FindOrCreateOwner(owner)
			if err != nil {
				return err
			}
			tID, err := FindOrCreateTeam(team)
			if err != nil {
				return err
			}

			toID, err := FindOrCreateTeamOwner(tID, oID)
			if err != nil {
				return err
			}
			log.Printf("Created Owner: %v (%v), Team: %v (%v), RelID: %v\n", owner, oID, team, tID, toID)
		}
	}

	return nil
}
