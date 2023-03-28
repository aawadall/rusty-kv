package persistence

import (
	"log"
	"os"
)

// PersistenceManager - persistence manager
type PersistenceManager struct {
	logger *log.Logger
	driver Driver
}

// NewPersistenceManager - create a new persistence manager
func NewPersistenceManager(config map[string]string) *PersistenceManager {
	pm := &PersistenceManager{
		logger: log.New(os.Stdout, "persistence: ", log.LstdFlags),
	}

	switch config["driver"] {
	case "flat_file":
		pm.driver = NewFlatFileDriver()
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
