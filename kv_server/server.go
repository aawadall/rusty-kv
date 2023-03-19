package kvserver

import (
	"log"

	"github.com/aawadall/simple-kv/types"
)

// TODO - Import packages here

// Aliases
type KVRecord = types.KVRecord

// KV Serever - A Struct that represents a KV Server
type KVServer struct {
	// TODO - Add fields here
	Records []KVRecord
	logger  *log.Logger
}

// NewKVServer - A function that creates a new KV Server
func NewKVServer() *KVServer {
	return &KVServer{
		Records: []KVRecord{},
		logger:  log.New(log.Writer(), "KVServer", log.LstdFlags), // TODO = Read from config
	}
}

// Start - A function that starts the KV Server
func (s *KVServer) Start() {
	// TODO - Start the KV Server here
	s.logger.Println("KV Server Started")
}

// Stop - A function that stops the KV Server
func (s *KVServer) Stop() {
	// TODO - Stop the KV Server here
	s.logger.Println("KV Server Stopped")
}
