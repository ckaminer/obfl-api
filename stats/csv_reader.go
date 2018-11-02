package stats

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadTeamOwners(path string) (map[string]string, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + path

	ownerToTeamMap := make(map[string]string)

	data, err := ReadCSV(filePath)
	if err != nil {
		return nil, err
	}

	for _, row := range data {
		ownerName := row[0]
		teamName := row[1]
		ownerToTeamMap[ownerName] = teamName
	}

	return ownerToTeamMap, nil
}

func ReadCSV(filePath string) ([][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	var data [][]string
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Unable to read file: %v due to: %v", filePath, err.Error())
		}
		if len(row) > 0 {
			data = append(data, row)
		}
	}

	return data, nil
}
