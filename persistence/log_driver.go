package persistence

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/aawadall/simple-kv/types"
)

// Mock Driver, writes Records to a flat file as log operations
type LogDriver struct {
	logFileName string
}

// NewLogDriver - create a new mock driver
func NewLogDriver(logFileName string) *LogDriver {
	driver := &LogDriver{
		logFileName: logFileName,
	}
	initFile(driver.logFileName)
	return driver
}

// implement Driver interface

// Write - write a record to disk
func (ff *LogDriver) Write(record KvRecord) error {
	// Open file for appending
	f, err := os.OpenFile(ff.logFileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write record to file in new line
	if _, err = f.WriteString(fmt.Sprintf("WRITE record(%v)\r", record.Key)); err != nil {
		return err
	}

	return nil
}

// Read - read a record from disk
func (ff *LogDriver) Read(key string) (KvRecord, error) {
	// Makeup a record
	blob := []byte("mock blob")
	valuesContainer := types.NewValuesContainer(blob)

	record := types.KVRecord{}
	record.Key = key
	record.Value = valuesContainer
	
	return record, nil
}

// Delete - delete a record from disk
func (ff *LogDriver) Delete(key string) error {
	// Open file for appending
	f, err := os.OpenFile(ff.logFileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write record to file in new line
	if _, err = f.WriteString(fmt.Sprintf("DELETE record(%v)\r", key)); err != nil {
		return err
	}

	return nil
}

// Compare - compare a record to disk
func (ff *LogDriver) Compare(record KvRecord) (bool, error) {
	// Random Comparison
	comparison := rand.Intn(2) == 0
	return comparison, nil
}

// Load - load all records from disk
func (ff *LogDriver) Load() ([]KvRecord, error) {
	// Return mock records
	return makeRecords(10), nil
}

// Helper functions
func makeRecords(count int) []KvRecord {
	records := make([]KvRecord, count)
	for i := 0; i < count; i++ {
		records[i] = *types.NewKVRecord(fmt.Sprintf("key%v", i), []byte(fmt.Sprintf("value%v", i)))
	}
	return records
}

func initFile(fileName string) {
	// Create file if it doesn't exist
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
