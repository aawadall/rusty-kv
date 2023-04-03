package persistence

// NoPersistence is a struct that implements the Persistence interface
// It is used when the user does not want to persist data

type NoPersistence struct {
}

// NewNoPersistence returns a new NoPersistence struct
func NewNoPersistence() *NoPersistence {
	return &NoPersistence{}
}

// Write(KvRecord) error
func (np *NoPersistence) Write(KvRecord) error {
	return nil
}

// Read(string) (KvRecord, error)
func (np *NoPersistence) Read(string) (KvRecord, error) {
	return KvRecord{}, nil
}

// Delete(string) error
func (np *NoPersistence) Delete(string) error {
	return nil
}

// Compare(KvRecord) (bool, error)
func (np *NoPersistence) Compare(KvRecord) (bool, error) {
	return false, nil
}

// Load() ([]KvRecord, error)
func (np *NoPersistence) Load() ([]KvRecord, error) {
	return []KvRecord{}, nil
}
