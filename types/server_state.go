package types

// Server State is an enum that represents the state of the server
type ServerState int

const (
	// ServerStarting - The server is starting
	ServerStarting ServerState = iota
	// ServerRunning - The server is running
	ServerRunning
	// ServerStopping - The server is stopping
	ServerStopping
	// ServerStopped - The server is stopped
	ServerStopped
	// ServerError - The server is in an error state
	ServerError
	// ServerUnknownState - The server is in an unknown state
	ServerUnknownState
)
