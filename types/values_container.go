package types

import (
	"fmt"
	"sync"
)

type ValuesContainer struct {
	mu    sync.Mutex
	Value [][]byte
}

func NewValuesContainer(value []byte) *ValuesContainer {
	container := &ValuesContainer{
		Value: make([][]byte, 0),
	}
	container.Set(value)
	return container
}

func (c *ValuesContainer) Get(version int) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// if version is out of range return nil
	if version > len(c.Value)-1 || version == 0 || version < -1 {
		return nil, fmt.Errorf("version %d is out of range", version)
	}

	// if version is -1 get last version
	if version < 0 {
		version = len(c.Value) - 1
	}

	return c.Value[version], nil
}

func (c *ValuesContainer) Set(value []byte) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value = append(c.Value, value)

	// get the version of the value
	version := len(c.Value) - 1
	return version
}

func (c *ValuesContainer) GetVersion() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.Value) - 1
}

func (c *ValuesContainer) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.Value)
}
