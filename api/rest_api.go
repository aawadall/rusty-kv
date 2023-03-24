package api

import (
	"log"
	"net/http"
	"os"

	kvserver "github.com/aawadall/simple-kv/kv_server"
)

// REST API for the application
type RestApi struct {
	logger *log.Logger
	server *kvserver.KVServer
}

// NewRestApi creates a new REST API
func NewRestApi(server *kvserver.KVServer) *RestApi {
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
	http.HandleFunc("/api/status", api.handleStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
