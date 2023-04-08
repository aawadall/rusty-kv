package kvserver

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aawadall/simple-kv/types"
)

// Test Server Creation
func TestNewKVServer(t *testing.T) {
	defer quiet()()
	// Arrange
	config := map[string]string{}
	// Act
	svr := NewKVServer(config)
	expectedType := "*kvserver.KVServer"

	// Assert that server is not nil
	if svr == nil {
		t.Errorf("Server is nil")
	}

	// Assert that server is of type KVServer
	svr_type := fmt.Sprintf("%T", svr)
	if svr_type != "*kvserver.KVServer" {
		t.Errorf("Server is of type %s instead of %s", svr_type, expectedType)
	}
}

// Test Server States
func TestKVServerStates(t *testing.T) {
	defer quiet()()
	// Arrange
	config := map[string]string{}
	svr := NewKVServer(config)

	// Assert that server is in the correct state
	if svr.state != types.ServerUnknownState {
		t.Errorf("server is in state %s instead of %s", stringState(svr.state), stringState(types.ServerUnknownState))
	}

	// Act
	svr.Start()

	// sleep for 17 seconds to allow server to start
	time.Sleep(12 * time.Second)

	// Assert that server is in the correct state
	if svr.state != types.ServerRunning {
		t.Errorf("server is in state %s instead of %s", stringState(svr.state), stringState(types.ServerRunning))
	}

	// Act
	svr.Stop()

	// sleep for 17 seconds to allow server to stop
	time.Sleep(12 * time.Second)

	// Assert that server is in the correct state
	if svr.state != types.ServerStopped {
		t.Errorf("server is in state %s instead of %s", stringState(svr.state), stringState(types.ServerStopped))
	}

}

// Helper function to convert state to string
func stringState(state types.ServerState) string {
	switch state {
	case types.ServerError:
		return "Error"
	case types.ServerRunning:
		return "Running"
	case types.ServerStarting:
		return "Starting"
	case types.ServerStopping:
		return "Stopping"
	case types.ServerStopped:
		return "Stopped"
	case types.ServerUnknownState:
		return "Unknown"
	default:
		return "Unknown"
	}
}

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	stdout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null

	return func() {
		defer null.Close()
		os.Stdout = stdout
		os.Stderr = serr
	}
}
