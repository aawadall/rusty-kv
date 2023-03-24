package api

import (
	"log"
	"net/http"
	"os"

	"github.com/aawadall/simple-kv/types"
)

// REST API for the application
type RestApi struct {
	logger *log.Logger
	server types.Server
}

// NewRestApi creates a new REST API
func NewRestApi(server types.Server) *RestApi {
	// TODO: Add configuration
	return &RestApi{
		logger: log.New(os.Stdout, "RestApi ", log.LstdFlags),
		server: server,
	}
}

// Start starts the REST API
func (api *RestApi) Start() {
	api.logger.Println("Starting REST API")
	go api.handleRequest()
}

// Stop stops the REST API
func (api *RestApi) Stop() {
	api.logger.Println("Stopping REST API")
}

// handle requests
func (api *RestApi) handleRequest() {
	// Server Router
	http.HandleFunc("/api/server/status", api.handleStatus)
	http.HandleFunc("/api/server/start", api.handleStart)
	http.HandleFunc("/api/server/stop", api.handleStop)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
