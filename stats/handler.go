package stats

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAllOwnersHandler(w http.ResponseWriter, r *http.Request) {
	owners, err := GetAllOwners()
	if err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, "Unable to get owners list")
	}

	RespondWithJSON(w, http.StatusOK, owners)
}