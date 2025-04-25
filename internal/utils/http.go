package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.Printf("❌ %s: %v", message, err)
	http.Error(w, message, statusCode)
}

func WriteJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Println("❌ Error encoding JSON response:", err)
	}
}

func DecodeJSONRequest[T any](w http.ResponseWriter, r *http.Request, dst *T) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		HandleError(w, err, "❌ Invalid JSON format", http.StatusBadRequest)
		return false
	}
	return true
}

func GetIDFromRequest(r *http.Request) (int, error) {
	return strconv.Atoi(mux.Vars(r)["id"])
}
