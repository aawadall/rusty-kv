package persistence

// FlatFileDriver - flat file driver

type FlatFileDriver struct {
}

// NewFlatFileDriver - create a new flat file driver
func NewFlatFileDriver() *FlatFileDriver {
	return &FlatFileDriver{}
}

// Write - write a record to disk
func (ff *FlatFileDriver) Write(record KvRecord) error {
	// TODO: implement
	return nil
}

// Read - read a record from disk
func (ff *FlatFileDriver) Read(key string) (KvRecord, error) {
	// TODO: implement
	return KvRecord{}, nil
}

// Delete - delete a record from disk
func (ff *FlatFileDriver) Delete(key string) error {
	// TODO: implement
	return nil
}

// Compare - compare a record to disk
func (ff *FlatFileDriver) Compare(record KvRecord) (bool, error) {
	// TODO: implement
	return false, nil
}

// Load - load all records from disk
func (ff *FlatFileDriver) Load() ([]KvRecord, error) {
	// TODO: implement
	return []KvRecord{}, nil
}
