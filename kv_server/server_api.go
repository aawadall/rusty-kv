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
	record, ok := s.Records.Get(key)
	if !ok {
		return nil, fmt.Errorf("key not found")
	}
	// otherwise return the value
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
	record, ok := s.Records.Get(key)
	if !ok {
		// if not, create a new record
		s.Records.Set(key, *types.NewKVRecord(key, bValue))
	} else {
		// otherwise update the value
		// Update the record
		record.UpdateRecord(key, bValue)

		s.Records.Set(key, record)
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
	_, ok := s.Records.Get(key)
	if !ok {
		return fmt.Errorf("key not found")
	}

	// otherwise delete the record
	s.Records.Delete(key)

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
	record, ok := s.Records.Get(key)
	if !ok {
		return fmt.Errorf("key not found")
	}

	// otherwise set the metadata
	newVer, err := record.SetMetadata(metadataKey, metadataValue)

	if err != nil {
		return err
	}

	// update version
	record.SetMetadata("Version", fmt.Sprintf("%d", newVer))

	// update the record
	s.Records.Set(key, record)
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
	record, ok := s.Records.Get(key)
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	// check if the metadata key is in the store
	if _, ok := record.Metadata[metadataKey]; !ok {
		return "", fmt.Errorf("metadata key not found")
	}

	// otherwise get the metadata
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
	record, ok := s.Records.Get(key)
	if !ok {
		return fmt.Errorf("key not found")
	}

	// check if the metadata key is in the store
	if _, ok := record.Metadata[metadataKey]; !ok {
		return fmt.Errorf("metadata key not found")
	}

	// otherwise delete the metadata
	record.DeleteMetadata(metadataKey)

	// update the record
	s.Records.Set(key, record)

	return nil
}

// Get All Metadata
func (s *KVServer) GetAllMetadata(key string) (metadata map[string]string, err error) {
	// check if the key is empty
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}

	// check if the key is in the store
	record, ok := s.Records.Get(key)
	if !ok {
		return nil, fmt.Errorf("key not found")
	}

	// otherwise get all metadata
	return record.ListMetadata()
}

// Find by partial key
func (s *KVServer) Find(partialKey string) (keys []string, err error) {
	// check if the partial key is empty
	if partialKey == "" {
		return nil, fmt.Errorf("partial key cannot be empty")
	}

	matchingKeys := s.Records.Find(partialKey)

	return matchingKeys, nil
}

// Find by Metadata and comparison operators
func (s *KVServer) FindByMetadata(query string) (keys []string, err error) {
	// Assuming query is commma separated entries, each in the format of "key:operator:value"
	// e.g. "name:contains:John,age:>=:18"
	// also assuming that queries are ANDed together

	// this is an expensive operation that is O(nxm) where n is the number of records and m is the number of query entries
	//  check if the query is empty
	if query == "" {
		return nil, fmt.Errorf("query cannot be empty")
	}

	// TODO - clean here
	// split the query into individual entries
	// queryEntries := strings.Split(query, ",")

	// records := s.Records

	// // loop through the entries
	// for _, entry := range queryEntries {
	// 	// split the entry into key, operator and value
	// 	entryParts := strings.Split(entry, ":")

	// 	// check if the entry is in the correct format
	// 	if len(entryParts) != 3 {
	// 		return nil, fmt.Errorf("invalid query entry '%s'", entry)
	// 	}

	// 	// refine the records
	// 	records, err = refineRecords(records, entryParts)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// }

	// // return the keys
	// keys = []string{}

	// for key := range records {
	// 	keys = append(keys, key)
	// }

	keys = s.Records.FindByMetadata(query)
	return keys, nil
}
