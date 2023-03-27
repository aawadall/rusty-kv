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
