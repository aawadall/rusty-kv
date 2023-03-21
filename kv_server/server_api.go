package kvserver

// Server API

// Get - A function that gets a value from the KV Server
func (s *KVServer) Get(key string) (value interface{}, err error) {
	// TODO - Implement this function
	return nil, nil
}

// Set - A function that sets a value in the KV Server
func (s *KVServer) Set(key string, value interface{}) (err error) {
	// TODO - Implement this function
	return nil
}

// Delete - A function that deletes a value from the KV Server
func (s *KVServer) Delete(key string) (err error) {
	// TODO - Implement this function
	return nil
}

// Advanced Methods
// Set Metadata
func (s *KVServer) SetMetadata(key string, value interface{}) (err error) {
	// TODO - Implement this function
	return nil
}

// Get Metadata
func (s *KVServer) GetMetadata(key string) (value interface{}, err error) {
	// TODO - Implement this function
	return nil, nil
}

// Delete Metadata
func (s *KVServer) DeleteMetadata(key string) (err error) {
	// TODO - Implement this function
	return nil
}

// Get All Metadata
func (s *KVServer) GetAllMetadata() (metadata map[string]interface{}, err error) {
	// TODO - Implement this function
	return nil, nil
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
