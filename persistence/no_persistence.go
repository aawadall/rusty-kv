package persistence

import (
	"log"
	"os"
)

// NoPersistence is a struct that implements the Persistence interface
// It is used when the user does not want to persist data

type NoPersistence struct {
	logger *log.Logger
}

// NewNoPersistence returns a new NoPersistence struct
func NewNoPersistence() *NoPersistence {
	return &NoPersistence{
		logger: log.New(os.Stdout, "persistence: ", log.LstdFlags),
	}
}

// Write(KvRecord) error
func (np *NoPersistence) Write(record KvRecord) error {
	np.logger.Printf("NoPersistence: Write() called for key: %v", record.Key)
	return nil
}

// Read(string) (KvRecord, error)
func (np *NoPersistence) Read(key string) (KvRecord, error) {
	np.logger.Printf("NoPersistence: Read() called for key: %v", key)
	return KvRecord{}, nil
}

// Delete(string) error
func (np *NoPersistence) Delete(key string) error {
	np.logger.Printf("NoPersistence: Delete() called for key: %v", key)
	return nil
}

// Compare(KvRecord) (bool, error)
func (np *NoPersistence) Compare(record KvRecord) (bool, error) {
	np.logger.Printf("NoPersistence: Compare() called for key: %v", record.Key)
	return false, nil
}

// Load() ([]KvRecord, error)
func (np *NoPersistence) Load() ([]KvRecord, error) {
	np.logger.Printf("NoPersistence: Load() called")
	return []KvRecord{}, nil
}
