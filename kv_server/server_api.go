package kvserver

import (
	"fmt"

	"github.com/aawadall/simple-kv/types"
)

// Server API

// Get - A function that gets a value from the KV Server
func (s *KVServer) Get(key string) (value interface{}, err error) {

	// check if the key is empty
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}

	// check if the key is in the store
	if _, ok := s.Records[key]; !ok {
		return nil, fmt.Errorf("key not found")
	}
	// otherwise return the value
	record := s.Records[key]
	value, err = record.GetValue(-1)
	return value, err
}

// Set - A function that sets a value in the KV Server
func (s *KVServer) Set(key string, value interface{}) (err error) {
	// check if the key is empty
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	// check if the value is empty
	if value == nil {
		return fmt.Errorf("value cannot be empty")
	}

	// cast value to bytes
	bValue := value.([]byte)

	// check if the key is in the store
	if _, ok := s.Records[key]; !ok {
		// if not, create a new record
		s.Records[key] = *types.NewKVRecord(key, bValue)
	} else {
		// otherwise update the value
		record := s.Records[key]
		record.UpdateRecord(key, bValue)
		s.Records[key] = record
	}

	// TODO - Implement this function
	return nil
}

// Delete - A function that deletes a value from the KV Server
func (s *KVServer) Delete(key string) (err error) {
	// check if the key is empty
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	// check if the key is in the store
	if _, ok := s.Records[key]; !ok {
		return fmt.Errorf("key not found")
	}

	// otherwise delete the record
	delete(s.Records, key)

	return nil
}

// Advanced Methods
// Set Metadata
func (s *KVServer) SetMetadata(key string, metadataKey string, metadataValue string) (err error) {
	// check if the key is empty
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	// check if the metadata key is empty
	if metadataKey == "" {
		return fmt.Errorf("metadata key cannot be empty")
	}

	// check if the metadata value is empty
	if metadataValue == "" {
		return fmt.Errorf("metadata value cannot be empty")
	}

	// check if the key is in the store
	if _, ok := s.Records[key]; !ok {
		return fmt.Errorf("key not found")
	}

	// otherwise set the metadata
	record := s.Records[key]
	record.SetMetadata(metadataKey, metadataValue)

	return nil
}

// Get Metadata
func (s *KVServer) GetMetadata(key string, metadataKey string) (value string, err error) {
	// check if the key is empty
	if key == "" {
		return "", fmt.Errorf("key cannot be empty")
	}

	// check if the metadata key is empty
	if metadataKey == "" {
		return "", fmt.Errorf("metadata key cannot be empty")
	}

	// check if the key is in the store
	if _, ok := s.Records[key]; !ok {
		return "", fmt.Errorf("key not found")
	}

	// check if the metadata key is in the store
	if _, ok := s.Records[key].Metadata[metadataKey]; !ok {
		return "", fmt.Errorf("metadata key not found")
	}

	// otherwise get the metadata
	record := s.Records[key]
	return record.GetMetadata(metadataKey)
}

// Delete Metadata
func (s *KVServer) DeleteMetadata(key string, metadataKey string) (err error) {
	// check if the key is empty
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	// check if the metadata key is empty
	if metadataKey == "" {
		return fmt.Errorf("metadata key cannot be empty")
	}

	// check if the key is in the store
	if _, ok := s.Records[key]; !ok {
		return fmt.Errorf("key not found")
	}

	// check if the metadata key is in the store
	if _, ok := s.Records[key].Metadata[metadataKey]; !ok {
		return fmt.Errorf("metadata key not found")
	}

	// otherwise delete the metadata
	record := s.Records[key]
	record.DeleteMetadata(metadataKey)

	return nil
}

// Get All Metadata
func (s *KVServer) GetAllMetadata(key string) (metadata map[string]string, err error) {
	// check if the key is empty
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}

	// check if the key is in the store
	if _, ok := s.Records[key]; !ok {
		return nil, fmt.Errorf("key not found")
	}

	// otherwise get all metadata
	record := s.Records[key]
	return record.ListMetadata()
}

// Find by partial key
func (s *KVServer) Find(partialKey string) (keys []string, err error) {
	// TODO - Implement this function
	return nil, nil
}

// Find by Metadata and comparison operators
func (s *KVServer) FindByMetadata(query string) (keys []string, err error) {
	// TODO - Implement this function
	return nil, nil
}
