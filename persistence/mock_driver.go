package persistence

import (
	"fmt"

	"log"
	"os"

	"github.com/aawadall/simple-kv/types"
)

// Mock Driver - Pretends to be a persistence driver
type MockDriver struct {
	records MockContainer
	logger  *log.Logger
}

// NewMockDriver - Create a new mock driver
func NewMockDriver() *MockDriver {
	driver := &MockDriver{
		records: *NewMockContainer(),
		logger:  log.New(os.Stdout, "MockDriver: ", log.LstdFlags),
	}
	driver.Populate(10)
	return driver
}

// implement Driver interface
// Write(KvRecord) error
func (md *MockDriver) Write(record KvRecord) error {
	md.logger.Printf("Write(%v)", record.Key)
	md.records.Set(record.Key, &record)
	return nil
}

// Read(string) (KvRecord, error)
func (md *MockDriver) Read(key string) (KvRecord, error) {
	md.logger.Printf("Read(%v)", key)
	record, ok := md.records.Get(key)
	if !ok {
		return record, fmt.Errorf("record not found")
	}
	return record, nil
}

// Delete(string) error
func (md *MockDriver) Delete(key string) error {
	md.logger.Printf("Delete(%v)", key)
	md.records.Delete(key)
	return nil
}

// Compare(KvRecord) (bool, error)
func (md *MockDriver) Compare(record KvRecord) (bool, error) {
	md.logger.Printf("Compare(%v)", record.Key)
	found, ok := md.records.Get(record.Key)
	if !ok {
		return false, fmt.Errorf("record not found")
	}
	return CompareRecords(record, found), nil
}

// Load() ([]KvRecord, error)
func (md *MockDriver) Load() ([]KvRecord, error) {
	md.logger.Printf("Load()")
	records := md.records.GetAll()
	return MapToRecordList(records), nil
}

// Populate - Populate the mock driver with some records
func (md *MockDriver) Populate(count int) {
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("key-%d", i)
		blob := []byte(fmt.Sprintf("value-%d", i))
		record := types.NewKVRecord(key, blob)
		md.records.Set(key, record)
	}
}

// Helper Functions
// MapToRecordList - Convert a map of records to a list of records
func MapToRecordList(records map[string]KvRecord) []KvRecord {
	var list []KvRecord
	for _, v := range records {
		list = append(list, v)
	}
	return list
}
