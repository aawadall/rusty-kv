package api

import (
	"log"
	"net/http"
	"os"

	"github.com/aawadall/simple-kv/types"
	"github.com/gorilla/mux"
)

// REST API for the application
type RestApi struct {
	logger *log.Logger
	server types.Server
	router *mux.Router
}

// NewRestApi creates a new REST API
func NewRestApi(server types.Server) *RestApi {
	// TODO: Add configuration
	return &RestApi{
		logger: log.New(os.Stdout, "RestApi ", log.LstdFlags),
		server: server,
		router: mux.NewRouter(),
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
	api.router.HandleFunc("/api/server/status", api.handleStatus)
	api.router.HandleFunc("/api/server/start", api.handleStart)
	api.router.HandleFunc("/api/server/stop", api.handleStop)

	// KV Router
	api.router.HandleFunc("/api/kv/{key}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			api.handleGet(w, r)
		case "POST":
			api.handleSet(w, r)
		case "DELETE":
			api.handleDelete(w, r)
		default:
			api.logger.Println("Invalid method")
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})

	// Metadata Router
	api.router.HandleFunc("/api/kv/{key}/metadata/{metadataKey}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			api.handleGetMetadata(w, r)
		case "POST":
			api.handleSetMetadata(w, r)
		case "DELETE":
			api.handleDeleteMetadata(w, r)
		default:
			api.logger.Println("Invalid method")
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})

	// Get All Metadata
	api.router.HandleFunc("/api/kv/{key}/metadata", api.handleGetAllMetadata)

	// Search Router
	api.router.HandleFunc("/api/kv/search/{partialKey}", api.handleFind)
	api.router.HandleFunc("/api/kv/search/metadata/{query}", api.handleFindByMetadata)
	log.Fatal(http.ListenAndServe(":8080", api.router))
}
