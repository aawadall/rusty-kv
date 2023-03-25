package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

// handler server start
func (api *RestApi) handleStart(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling start request")
	// Start server
	api.server.Start()

	// Write status to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Server started")
}

// handler server stop
func (api *RestApi) handleStop(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling stop request")
	// Stop server
	api.server.Stop()

	// Write status to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Server stopped")
}

// handle Get(key string) (interface{}, error)
func (api *RestApi) handleGet(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling get request")
	// Get key from request
	vars := mux.Vars(r)

	key, ok := vars["key"]
	if !ok {
		api.logger.Printf("Something went wrong %v, do we have key? %v", ok, key)

		// api.logger.Println("No key provided")
		// http.Error(w, "No key provided", http.StatusBadRequest)
		// return
	}

	if key == "" {
		api.logger.Println("No key provided")
		http.Error(w, "No key provided", http.StatusBadRequest)
		return
	}

	// Get valueBytes from server
	valueBytes, err := api.server.Get(key)

	if err != nil {
		api.logger.Println("Error getting value from server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// cast valueBytes to byte array
	var value []byte
	if valueBytes != nil {
		value = valueBytes.([]byte)
	}

	// convert value to string
	valueString := string(value)

	// Write value to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(valueString)
}

// handle Set(key string, value interface{}) error
func (api *RestApi) handleSet(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling set request")
	// Get key from request

	vars := mux.Vars(r)
	key, ok := vars["key"]
	if !ok {
		api.logger.Printf("Something went wrong %v, do we have key? %v", ok, key)

		// api.logger.Println("No key provided")
		// http.Error(w, "No key provided", http.StatusBadRequest)
		// return
	}

	if key == "" {
		api.logger.Println("No key provided")
		http.Error(w, "No key provided", http.StatusBadRequest)
		return
	}

	// Get value from request
	value := r.Body
	if value == nil {
		api.logger.Println("No value provided")
		http.Error(w, "No value provided", http.StatusBadRequest)
		return
	}

	// convert value to byte array
	buf := new(bytes.Buffer)
	buf.ReadFrom(value)
	valueBytes := buf.Bytes()

	// Set value in server
	err := api.server.Set(key, valueBytes)
	if err != nil {
		api.logger.Println("Error setting value in server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write status to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Value set")
}

// handle Delete(key string) error
func (api *RestApi) handleDelete(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling delete request")
	// Get key from request
	vars := mux.Vars(r)
	key, ok := vars["key"]
	if !ok {
		api.logger.Println("No key provided")
		http.Error(w, "No key provided", http.StatusBadRequest)
		return
	}

	if key == "" {
		api.logger.Println("No key provided")
		http.Error(w, "No key provided", http.StatusBadRequest)
		return
	}

	// Delete value from server
	err := api.server.Delete(key)
	if err != nil {
		api.logger.Println("Error deleting value from server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write status to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Value deleted")
}

// handle SetMetadata(key string, metadataKey string, metadataValue string) error
func (api *RestApi) handleSetMetadata(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling set metadata request")
	// Get key from request
	key := r.URL.Query().Get("key")
	if key == "" {
		api.logger.Println("No key provided")
		http.Error(w, "No key provided", http.StatusBadRequest)
		return
	}

	// Get metadata key from request
	metadataKey := r.URL.Query().Get("metadataKey")
	if metadataKey == "" {
		api.logger.Println("No metadata key provided")
		http.Error(w, "No metadata key provided", http.StatusBadRequest)
		return
	}

	// Get metadata value from request
	metadataValue := r.URL.Query().Get("metadataValue")
	if metadataValue == "" {
		api.logger.Println("No metadata value provided")
		http.Error(w, "No metadata value provided", http.StatusBadRequest)
		return
	}

	// Set metadata in server
	err := api.server.SetMetadata(key, metadataKey, metadataValue)
	if err != nil {
		api.logger.Println("Error setting metadata in server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write status to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Metadata set")
}

// handle GetMetadata(key string, metadataKey string) (string, error)
func (api *RestApi) handleGetMetadata(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling get metadata request")
	// Get key from request
	key := r.URL.Query().Get("key")
	if key == "" {
		api.logger.Println("No key provided")
		http.Error(w, "No key provided", http.StatusBadRequest)
		return
	}

	// Get metadata key from request
	metadataKey := r.URL.Query().Get("metadataKey")
	if metadataKey == "" {
		api.logger.Println("No metadata key provided")
		http.Error(w, "No metadata key provided", http.StatusBadRequest)
		return
	}

	// Get metadata from server
	metadata, err := api.server.GetMetadata(key, metadataKey)
	if err != nil {
		api.logger.Println("Error getting metadata from server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write metadata to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(metadata)
}

// handle DeleteMetadata(key string, metadataKey string) error
func (api *RestApi) handleDeleteMetadata(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling delete metadata request")
	// Get key from request
	key := r.URL.Query().Get("key")
	if key == "" {
		api.logger.Println("No key provided")
		http.Error(w, "No key provided", http.StatusBadRequest)
		return
	}

	// Get metadata key from request
	metadataKey := r.URL.Query().Get("metadataKey")
	if metadataKey == "" {
		api.logger.Println("No metadata key provided")
		http.Error(w, "No metadata key provided", http.StatusBadRequest)
		return
	}

	// Delete metadata from server
	err := api.server.DeleteMetadata(key, metadataKey)
	if err != nil {
		api.logger.Println("Error deleting metadata from server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write status to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Metadata deleted")
}

// handle GetAllMetadata(key string) (map[string]string, error)
func (api *RestApi) handleGetAllMetadata(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling get all metadata request")
	// Get key from request
	key := r.URL.Query().Get("key")
	if key == "" {
		api.logger.Println("No key provided")
		http.Error(w, "No key provided", http.StatusBadRequest)
		return
	}

	// Get all metadata from server
	metadata, err := api.server.GetAllMetadata(key)
	if err != nil {
		api.logger.Println("Error getting all metadata from server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write metadata to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(metadata)
}

// handle Find(partialKey string) ([]string, error)
func (api *RestApi) handleFind(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling find request")
	// Get partial key from request
	partialKey := r.URL.Query().Get("partialKey")
	if partialKey == "" {
		api.logger.Println("No partial key provided")
		http.Error(w, "No partial key provided", http.StatusBadRequest)
		return
	}

	// Find keys from server
	keys, err := api.server.Find(partialKey)
	if err != nil {
		api.logger.Println("Error finding keys from server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write keys to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(keys)
}

// handle FindByMetadata(query string) ([]string, error)
func (api *RestApi) handleFindByMetadata(w http.ResponseWriter, r *http.Request) {
	api.logger.Println("Handling find by metadata request")
	// Get query from request
	query := r.URL.Query().Get("query")
	if query == "" {
		api.logger.Println("No query provided")
		http.Error(w, "No query provided", http.StatusBadRequest)
		return
	}

	// Find keys from server
	keys, err := api.server.FindByMetadata(query)
	if err != nil {
		api.logger.Println("Error finding keys from server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write keys to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(keys)
}
