package types

import "sync"

type MetadataContainer struct {
	mu       sync.Mutex
	Metadata map[string]string
}

func NewMetadataContainer() *MetadataContainer {
	return &MetadataContainer{
		Metadata: make(map[string]string),
	}
}

func (c *MetadataContainer) Set(key string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Metadata[key] = value
}

func (c *MetadataContainer) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.Metadata[key]
	return value, ok
}

func (c *MetadataContainer) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Metadata, key)
}

func (c *MetadataContainer) GetAll() map[string]string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Metadata
}
