package kvserver

import (
	"log"
	"sync"
	"time"

	"github.com/aawadall/simple-kv/api"
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
	rest    *api.RestApi
}

// NewKVServer - A function that creates a new KV Server
func NewKVServer() *KVServer {
	server := &KVServer{
		Records: make(map[string]KVRecord),
		logger:  log.New(log.Writer(), "KVServer", log.LstdFlags),
		config:  config.NewConfigurationManager(),
		state:   types.ServerUnknownState,
	}
	server.rest = api.NewRestApi(server)

	return server
}

// Start - A function that starts the KV Server
func (s *KVServer) Start() {
	// TODO - Start the KV Server here
	wg := &sync.WaitGroup{}

	wg.Add(1)
	// Event loop in a goroutine
	go func() {
		s.state = types.ServerStarting

		// Start the REST API
		s.rest.Start()

		// TODO - Start the KV Server here

		s.state = types.ServerRunning
		for s.state != types.ServerStopped &&
			s.state != types.ServerError {
			// TODO - Event loop code here
			time.Sleep(1 * time.Second)

			// translate state to string
			//state := stateToString(s)
			// Write to log

			//s.logger.Printf("KV Server is %s", state)
		}
		wg.Done()
	}()
	wg.Wait()
}

func stateToString(s *KVServer) string {
	state := ""

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
	return state
}

// Stop - A function that stops the KV Server
func (s *KVServer) Stop() {
	// TODO - Stop the KV Server here
	s.logger.Println("KV Server Stopping")
	s.state = types.ServerStopping
	// TODO - Stop the KV Server here
	s.rest.Stop()
	// Sleep for a bit to simulate real stopping
	time.Sleep(5 * time.Second)

	// Write to log
	s.state = types.ServerStopped
	s.logger.Println("KV Server Stopped")
}

// GetStatus - A function that returns the status of the KV Server
func (s *KVServer) GetStatus() (interface{}, error) {
	// TODO - Return the status of the KV Server here
	status := make(map[string]string)
	status["state"] = stateToString(s)
	return status, nil
}
