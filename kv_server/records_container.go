package kvserver

import (
	"fmt"
	"strings"
	"sync"
)

type Container struct {
	mu      sync.Mutex
	Records map[string]KVRecord
}

func NewContainer() *Container {
	return &Container{
		Records: make(map[string]KVRecord),
	}
}

func (c *Container) Get(key string) (KVRecord, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record, ok := c.Records[key]
	return record, ok
}

func (c *Container) Set(key string, record KVRecord) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Records[key] = record
}

func (c *Container) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Records, key)
}

func (c *Container) Find(partialKey string) []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	var keys []string
	for key := range c.Records {
		if partialKey == key[:len(partialKey)] {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c *Container) FindByMetadata(query string) []string {
	// TODO: Review
	c.mu.Lock()
	defer c.mu.Unlock()
	var keys []string
	for key, record := range c.Records {
		if _, found := record.Metadata.Get(query); found {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c *Container) GetMetadata(key string, metadataKey string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record, ok := c.Records[key]
	if !ok {
		return "", false
	}
	metadata, ok := record.Metadata.Get(metadataKey)
	return metadata, ok
}

func (c *Container) SetMetadata(key string, metadataKey string, metadataValue string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record := c.Records[key]
	record.Metadata.Set(metadataKey, metadataValue)
	c.Records[key] = record
}

func (c *Container) DeleteMetadata(key string, metadataKey string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record := c.Records[key]
	record.Metadata.Delete(metadataKey)
	c.Records[key] = record
}

func (c *Container) GetAllMetadata(key string) map[string]string {
	c.mu.Lock()
	defer c.mu.Unlock()
	record := c.Records[key]
	return record.Metadata.GetAll()
}

func (c *Container) List() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	var keys []string
	for key := range c.Records {
		keys = append(keys, key)
	}
	return keys
}

// Helper Functions
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func refineRecords(records map[string]KVRecord, entryParts []string) (refinedRecords map[string]KVRecord, err error) {
	// check if the entry is in the correct format
	if len(entryParts) != 3 {
		return nil, fmt.Errorf("invalid query entry '%s'", entryParts)
	}

	// get the key, operator and value
	 key := entryParts[0]
	operator := entryParts[1]
	value := entryParts[2]

	// loop through the records
	for recordKey, record := range records {
		// check if the key exists
		if _, found := record.Metadata.Get(key); !found {
			delete(records, recordKey)
			continue
		}

		// check if the operator is valid
		if !isValidOperator(operator) {
			return nil, fmt.Errorf("invalid operator '%s'", operator)
		}

		// check if the value matches
		metadata, _ := record.Metadata.Get(key)
		if !matches(metadata, operator, value) {
			delete(records, recordKey)
			continue
		}
	}

	return records, nil
}

func isValidOperator(operator string) bool {
	switch operator {
	case ">", ">=", "<", "<=", "==", "!=", "contains":
		return true
	default:
		return false
	}
}

func matches(value string, operator string, target string) bool {
	switch operator {
	case ">":
		return value > target
	case ">=":
		return value >= target
	case "<":
		return value < target
	case "<=":
		return value <= target
	case "==":
		return value == target
	case "!=":
		return value != target
	case "contains":
		return contains(value, target)
	default:
		return false
	}
}
