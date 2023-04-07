package persistence

import (
	"sync"

	"github.com/aawadall/simple-kv/types"
)

// Mock Container - A container for mock records
type MockContainer struct {
	mu   sync.Mutex
	data map[string]types.KVRecord
}

// NewMockContainer - Create a new mock container
func NewMockContainer() *MockContainer {
	return &MockContainer{
		data: make(map[string]types.KVRecord),
	}
}

// Get - Get a record from the container
func (c *MockContainer) Get(key string) (types.KVRecord, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record, ok := c.data[key]
	return record, ok
}

// Set - Set a record in the container
func (c *MockContainer) Set(key string, record *types.KVRecord) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = *record
}

// Delete - Delete a record from the container
func (c *MockContainer) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// GetAll - Get all records from the container
func (c *MockContainer) GetAll() map[string]types.KVRecord {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.data
}
