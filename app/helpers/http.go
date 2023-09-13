package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Type  string
	Value error
}
type Result struct {
	ErrorStatus  bool        `json:"errorStatus"`
	ErrorMessage *string     `json:"errorMessage"`
	Results      interface{} `json:"results"`
}

func RespondSuccess(w http.ResponseWriter, code int, result interface{}) {
	RespondWithJSON(w, code, Result{ErrorStatus: false, ErrorMessage: nil, Results: result})
}

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, Result{ErrorStatus: true, ErrorMessage: &message, Results: nil})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func VerifyErrors(w http.ResponseWriter, err Error) {
	log.Println(err)

	switch err.Type {
	case ErrorInternal:
		RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Internal Error.",
		})
	}
}
