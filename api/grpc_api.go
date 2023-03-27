package api

import (
	"log"
	"os"

	"github.com/aawadall/simple-kv/types"
)

// gRPC API for the application
type GrpcApi struct {
	logger *log.Logger
	server types.Server
}

// NewGrpcApi creates a new gRPC API
func NewGrpcApi(server types.Server) *GrpcApi {
	// TODO: Add configuration
	return &GrpcApi{
		logger: log.New(os.Stdout, "GrpcApi ", log.LstdFlags),
		server: server,
	}
}

// Start starts the gRPC API
func (api *GrpcApi) Start() {
	api.logger.Println("Starting gRPC API")
}

// Stop stops the gRPC API
func (api *GrpcApi) Stop() {
	api.logger.Println("Stopping gRPC API")
}
