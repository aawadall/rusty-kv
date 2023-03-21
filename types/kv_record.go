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
	Value    [][]byte
	Metadata map[string]string
}

// NewKVRecord - A function that creates a new KV Record
func NewKVRecord(key string, value []byte) *KVRecord {
	record := &KVRecord{
		Id:       uuid.New(),
		Key:      key,
		Value:    make([][]byte, 0),
		Metadata: make(map[string]string),
	}
	record.Value = append(record.Value, value)
	record.Metadata["Version"] = "1"
	return record
}

// Set Metadata - A function that sets the metadata for a KV Record
func (r *KVRecord) SetMetadata(key string, value string) (version int, err error) {
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
		r.Metadata = make(map[string]string)
		r.Metadata["Version"] = "1"
	}

	// if the key is already in the metadata, advance the version by 1
	version = r.GetVersion()

	if _, ok := r.Metadata[key]; ok {
		version = version + 1
		r.Metadata["Version"] = fmt.Sprintf("%d", version)

	}

	r.Metadata[key] = value

	return version, nil
}

// GetVersion - A function that gets the version of a KV Record
func (r *KVRecord) GetVersion() int {
	version, err := strconv.Atoi(r.Metadata["Version"])
	if err != nil {
		return 0
	}

	return version
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

	// if the key is not in the metadata, return error
	if _, ok := r.Metadata[key]; !ok {
		return "", fmt.Errorf("metadata `Key` not found")
	}

	return r.Metadata[key], nil
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

	// if the key is not in the metadata, return error
	if _, ok := r.Metadata[key]; !ok {
		return -1, fmt.Errorf("metadata `Key` not found")
	}

	// if the key is in the metadata, advance the version by 1
	version = r.GetVersion()
	version = version + 1
	r.Metadata["Version"] = fmt.Sprintf("%d", version)

	delete(r.Metadata, key)

	return version, nil
}

// List Metadata - A function that lists the metadata for a KV Record
func (r *KVRecord) ListMetadata() (metadata map[string]string, err error) {
	// if the metadata is nil, return error
	if r.Metadata == nil {
		return nil, fmt.Errorf("metadata is nil")
	}

	return r.Metadata, nil
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

	// if the metadata is nil, create it
	if r.Metadata == nil {
		r.Metadata = make(map[string]string)
		r.Metadata["Version"] = "1"
	}

	// if the key is already in the metadata, advance the version by 1
	version = r.GetVersion()

	version = version + 1
	r.Metadata["Version"] = fmt.Sprintf("%d", version)

	r.Key = key
	// append the value to the value array
	r.Value = append(r.Value, value)

	return version, nil
}

// Get Value at version
func (r *KVRecord) GetValue(version int) (value []byte, err error) {
	// if version is -1 return last version
	if version == -1 {
		return r.Value[len(r.Value)-1], nil
	}

	// if version is greater than the length of the value array return error
	if version > len(r.Value) {
		return nil, fmt.Errorf("version %d does not exist", version)
	}

	// if version is 0 return error
	if version == 0 {
		return nil, fmt.Errorf("version cannot be 0")
	}

	// if version is a positive number return the value at that index
	return r.Value[version-1], nil
}
