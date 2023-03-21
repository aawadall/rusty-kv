package kvserver

import (
	"log"
	"time"

	"github.com/aawadall/simple-kv/config"
	"github.com/aawadall/simple-kv/types"
)

// Aliases
type KVRecord = types.KVRecord
type ServerState = types.ServerState

// KV Serever - A Struct that represents a KV Server
type KVServer struct {
	// TODO - Add fields here
	Records map[string]KVRecord
	logger  *log.Logger
	state   ServerState
	config  *config.ConfigurationManager
}

// NewKVServer - A function that creates a new KV Server
func NewKVServer() *KVServer {
	return &KVServer{
		Records: make(map[string]KVRecord),
		logger:  log.New(log.Writer(), "KVServer", log.LstdFlags), // TODO = Read from config
		config:  config.NewConfigurationManager(),
		state:   types.ServerUnknownState,
	}
}

// Start - A function that starts the KV Server
func (s *KVServer) Start() {
	// TODO - Start the KV Server here

	// Event loop in a goroutine
	go func() {
		s.state = types.ServerStarting

		// TODO - Start the KV Server here
		// Sleep for a bit to simulate real starting
		time.Sleep(1 * time.Second)

		for s.state != types.ServerStopped &&
			s.state != types.ServerError {
			// TODO - Event loop code here
			time.Sleep(1 * time.Second)

			state := ""
			// translate state to string
			switch s.state {
			case types.ServerError:
				state = "Error"
			case types.ServerRunning:
				state = "Running"
			case types.ServerStarting:
				state = "Starting"
			case types.ServerStopping:
				state = "Stopping"
			case types.ServerStopped:
				state = "Stopped"
			case types.ServerUnknownState:
				state = "Unknown"

			}
			// Write to log

			s.logger.Printf("KV Server is %s", state)
		}
	}()
}

// Stop - A function that stops the KV Server
func (s *KVServer) Stop() {
	// TODO - Stop the KV Server here
	s.logger.Println("KV Server Stopping")
	s.state = types.ServerStopping
	// TODO - Stop the KV Server here
	// Sleep for a bit to simulate real stopping
	time.Sleep(10 * time.Second)

	// Write to log
	s.state = types.ServerStopped
	s.logger.Println("KV Server Stopped")
}
