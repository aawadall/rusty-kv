package persistence

// SQLiteDatabaseDriver - sqlite database driver

type SQLiteDatabaseDriver struct {
}

// NewSQLiteDatabaseDriver - create a new sqlite database driver
func NewSQLiteDatabaseDriver() *SQLiteDatabaseDriver {
	return &SQLiteDatabaseDriver{}
}

// Write - write a record to disk
func (ff *SQLiteDatabaseDriver) Write(record KvRecord) error {
	// TODO: implement
	return nil
}

// Read - read a record from disk
func (ff *SQLiteDatabaseDriver) Read(key string) (KvRecord, error) {
	// TODO: implement
	return KvRecord{}, nil
}

// Compare - compare a record to disk
func (ff *SQLiteDatabaseDriver) Compare(record KvRecord) (bool, error) {
	// TODO: implement
	return false, nil
}

// Load - load all records from disk
func (ff *SQLiteDatabaseDriver) Load() ([]KvRecord, error) {
	// TODO: implement
	return []KvRecord{}, nil
}
