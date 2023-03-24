package api

import (
	"encoding/json"
	"net/http"
)

// handle status
func (api *RestApi) handleStatus(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling status request")
	// Get status from server
	status, err := api.server.GetStatus()
	if err != nil {
		api.logger.Println("Error getting status from server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write status to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
