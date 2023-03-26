package types

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

// A Struct that represents a KV Record
type KVRecord struct {
	// define uuid id
	Id       uuid.UUID
	Key      string
	Value    *ValuesContainer
	Metadata *MetadataContainer
}

// NewKVRecord - A function that creates a new KV Record
func NewKVRecord(key string, value []byte) *KVRecord {
	record := &KVRecord{
		Id:       uuid.New(),
		Key:      key,
		Value:    NewValuesContainer(),
		Metadata: NewMetadataContainer(),
	}

	record.Metadata.Set("Version", "1")
	return record
}

// Set Metadata - A function that sets the metadata for a KV Record
func (r *KVRecord) SetMetadata(key string, value string) (int, error) {
	// if the key is empty return error
	if key == "" {
		return -1, fmt.Errorf("metadata `Key` cannot be empty")
	}

	// if the value is empty return error
	if value == "" {
		return -1, fmt.Errorf("metadata `Value` cannot be empty")
	}

	// if the metadata is nil, create it
	if r.Metadata == nil {
		r.Metadata = NewMetadataContainer()
	}

	r.Metadata.Set(key, value)

	// if the key is already in the metadata, advance the version by 1
	version := r.Value.GetVersion()

	return version, nil
}

// GetVersion - A function that gets the version of a KV Record
func (r *KVRecord) GetVersion() int {
	return r.Value.GetVersion()
}

// Get Metadata - A function that gets the metadata for a KV Record
func (r *KVRecord) GetMetadata(key string) (value string, err error) {
	// if the key is empty return error
	if key == "" {
		return "", fmt.Errorf("metadata `Key` cannot be empty")
	}

	// if the metadata is nil, return error
	if r.Metadata == nil {
		return "", fmt.Errorf("metadata is nil")
	}

	// // if the key is not in the metadata, return error
	// if _, ok := r.Metadata[key]; !ok {
	// 	return "", fmt.Errorf("metadata `Key` not found")
	// }
	metadata, ok := r.Metadata.Get(key)
	if !ok {
		return "", fmt.Errorf("metadata `Key` not found")
	}
	return metadata, nil
}

// Delete Metadata - A function that deletes the metadata for a KV Record
func (r *KVRecord) DeleteMetadata(key string) (version int, err error) {
	// if the key is empty return error
	if key == "" {
		return -1, fmt.Errorf("metadata `Key` cannot be empty")
	}

	// if the metadata is nil, return error
	if r.Metadata == nil {
		return -1, fmt.Errorf("metadata is nil")
	}

	r.Metadata.Delete(key)

	version = r.GetVersion()

	r.Metadata.Set("Version", strconv.Itoa(version))

	return version, nil
}

// List Metadata - A function that lists the metadata for a KV Record
func (r *KVRecord) ListMetadata() (metadata map[string]string, err error) {
	// if the metadata is nil, return error

	return r.Metadata.GetAll(), nil
}

// Update Record
func (r *KVRecord) UpdateRecord(key string, value []byte) (version int, err error) {

	// if the key is empty return error
	if key == "" {
		return -1, fmt.Errorf("metadata `Key` cannot be empty")
	}

	// if the value is empty return error
	if value == nil {
		return -1, fmt.Errorf("metadata `Value` cannot be empty")
	}

	// if the key is already in the metadata, advance the version by 1

	r.Key = key
	// append the value to the value array
	r.Value.Set(value)
	version = r.GetVersion()
	r.Metadata.Set("Version", strconv.Itoa(version))
	return version, nil
}

// Get Value at version
func (r *KVRecord) GetValue(version int) (value []byte, err error) {
	// if version is -1 return last version

	// if version is a positive number return the value at that index
	return r.Value.Get(version)
}
