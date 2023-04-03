package persistence

import (
	"fmt"
	"log"
	"os"
)

// PersistenceManager - persistence manager
type PersistenceManager struct {
	logger *log.Logger
	driver Driver
}

// NewPersistenceManager - create a new persistence manager
func NewPersistenceManager(config map[string]interface{}) *PersistenceManager {
	pm := &PersistenceManager{
		logger: log.New(os.Stdout, "persistence: ", log.LstdFlags),
	}

	pm.logger.Printf("Creating Persistence Manager with driver: %v", config["driver"])
	switch config["driver"] {
	case "flat_file":
		pm.driver = NewFlatFileDriver()
	case "sqlite":
		pm.driver = NewSQLiteDatabaseDriver(fmt.Sprintf("%v", config["db_location"]))
	case "mock":
		pm.driver = NewMockDriver(fmt.Sprintf("%v", config["file_location"]))
	case "none":
		pm.driver = NewNoPersistence()
	default:
		pm.driver = NewFlatFileDriver()
	}

	return pm
}

// Start - start the persistence manager
func (pm *PersistenceManager) Start() {
	// TODO: implement
	pm.logger.Println("Starting Persistence Manager")
}

// Stop - stop the persistence manager
func (pm *PersistenceManager) Stop() {
	// TODO: implement
	pm.logger.Println("Stopping Persistence Manager")
}

// Write - write a record to disk
func (pm *PersistenceManager) Write(record KvRecord) error {
	return pm.driver.Write(record)
}

// Read - read a record from disk
func (pm *PersistenceManager) Read(key string) (KvRecord, error) {
	return pm.driver.Read(key)
}

// Delete - delete a record from disk
func (pm *PersistenceManager) Delete(key string) error {
	return pm.driver.Delete(key)
}

// Compare - compare a record to disk
func (pm *PersistenceManager) Compare(record KvRecord) (bool, error) {
	return pm.driver.Compare(record)
}

// Load - load all records from disk
func (pm *PersistenceManager) Load() ([]KvRecord, error) {
	return pm.driver.Load()
}

// Save - save all records to disk
func (pm *PersistenceManager) Save(records []KvRecord) error {
	pm.logger.Printf("Saving %v records to disk", len(records))
	for _, record := range records {
		pm.logger.Print(".")
		err := pm.driver.Write(record)
		if err != nil {
			return err
		}
	}
	pm.logger.Printf("Done saving %v records to disk", len(records))
	return nil
}

// Sync - sync all records to disk
func (pm *PersistenceManager) Sync(records []KvRecord) error {
	pm.logger.Println("Syncing records to disk")

	// 1. Load all records from disk
	diskRecords, err := pm.driver.Load()
	if err != nil {
		return err
	}

	// 2. Compare With records, find any records in diskRecords and not in records
	delta := []string{}
	for _, diskRecord := range diskRecords {
		found := false
		for _, record := range records {
			if diskRecord.Key == record.Key {
				found = true
				break
			}
		}
		if !found {
			delta = append(delta, diskRecord.Key)
		}
	}

	// 3. Delete any records that are not in With records
	for _, key := range delta {
		pm.logger.Printf("Deleting record: %v", key)
		err := pm.driver.Delete(key)
		if err != nil {
			return err
		}
	}

	// 4. Write All records from records
	for _, record := range records {
		pm.logger.Printf("Writing record: %v", record.Key)
		err := pm.driver.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}
