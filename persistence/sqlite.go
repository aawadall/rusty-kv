package persistence

// SQLiteDatabaseDriver - sqlite database driver

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDatabaseDriver struct {
	dbLocation string
	db         *sql.DB
	logger     *log.Logger
}

// NewSQLiteDatabaseDriver - create a new sqlite database driver
func NewSQLiteDatabaseDriver(dbLocation string) *SQLiteDatabaseDriver {
	driver := &SQLiteDatabaseDriver{
		dbLocation: dbLocation,
		logger:     log.New(os.Stdout, "sqlite: ", log.LstdFlags),
	}

	driver.logger.Printf("Creating SQLite Database Driver with location: %v", dbLocation)
	// initialize the driver
	driver.init()

	return driver
}

// init - initialize the sqlite database driver
func (ff *SQLiteDatabaseDriver) init() {
	ff.logger.Printf("Initializing SQLite Database Driver with location: %v", ff.dbLocation)
	// TODO: implement
	var file *os.File
	// 1. Check if the database exists
	if _, err := os.Stat(ff.dbLocation); os.IsNotExist(err) {
		// 1.1. If it does not exist, create it
		file, err = os.Create(ff.dbLocation)
		if err != nil {
			panic(err)
		}
		file.Close()
	}
	// 2. Ensure our table exists
	ff.initDatabase()

}

// Write - write a record to disk
func (ff *SQLiteDatabaseDriver) Write(record KvRecord) error {
	// TODO: implement
	ff.logger.Printf("Writing record to SQLite Database Driver with key: %v", record.Key)
	// 1. Insert the record
	err := ff.insertRecord(record)
	if err != nil {
		return err
	}
	// 2. Insert the metadata
	err = ff.insertMetadata(record)
	if err != nil {
		return err
	}
	ff.logger.Printf("Wrote record to SQLite Database Driver with key: %v", record.Key)
	return nil
}

// Read - read a record from disk
func (ff *SQLiteDatabaseDriver) Read(key string) (KvRecord, error) {
	// TODO: implement
	ff.logger.Printf("Reading record from SQLite Database Driver with key: %v", key)
	if ff.findRecord(key) {
		// 1. Get the record
		record, err := ff.getRecord(key)
		if err != nil {
			return KvRecord{}, err
		}
		// 2. Get the metadata
		metadata, err := ff.getMetadata(key)
		if err != nil {
			return KvRecord{}, err
		}

		// Set the metadata
		for k, v := range metadata {
			record.Metadata.Set(k, v)
		}

		return record, nil
	}
	return KvRecord{}, nil
}

// Delete - delete a record from disk
func (ff *SQLiteDatabaseDriver) Delete(key string) error {
	// TODO: implement
	// 1. Delete they metadata
	err := ff.deleteMetadata(key)
	if err != nil {
		return err
	}

	// 2. Delete the record
	err = ff.deleteRecord(key)
	if err != nil {
		return err
	}

	return nil
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

// initDatabase - initialize the database
func (ff *SQLiteDatabaseDriver) initDatabase() {
	ff.db, _ = sql.Open("sqlite3", ff.dbLocation)

	// ensure record table exists
	// RECORDS TABLE
	// id - int (auto increment)
	// key - string (unique)
	// value - blob - latest value
	query := `CREATE TABLE IF NOT EXISTS records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT UNIQUE,
		value BLOB
	);`
	ff.db.Exec(query)

	// ensure metadata table exists
	// METADATA TABLE
	// id - int (auto increment)
	// key - string - FK to records table
	// metadataKey - string
	// metadataValue - string
	// UNIQUE (key, metadataKey)
	query = `CREATE TABLE IF NOT EXISTS metadata (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT,
		metadataKey TEXT,
		metadataValue TEXT,
		FOREIGN KEY(key) REFERENCES records(key),
		UNIQUE (key, metadataKey)
	);`
	ff.db.Exec(query)

	// ensure old values table exists
	// OLD VALUES TABLE
	// id - int (auto increment)
	// key - string - FK to records table
	// version - int
	// value - blob
	// UNIQUE (key, version)
	query = `CREATE TABLE IF NOT EXISTS oldValues (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT,
		version INTEGER,
		value BLOB,
		FOREIGN KEY(key) REFERENCES records(key),
		UNIQUE (key, version)
	);`
	ff.db.Exec(query)
}

// insertRecord - insert a record into the database
func (ff *SQLiteDatabaseDriver) insertRecord(record KvRecord) error {
	key := record.Key
	currentValue := record.Value
	version := record.GetVersion()

	ff.logger.Printf("Inserting record into SQLite Database Driver with key: %v", key)
	// upsert the record
	query := `INSERT OR REPLACE INTO records (key, value) VALUES (?, ?);`

	// TODO - remove after debugging
	ff.logger.Printf("Query: %v", query)
	ff.logger.Printf("Key: %v", key)
	ff.logger.Printf("Value: %v", currentValue)
	ff.logger.Printf("Value Type: %T", currentValue)

	result, err := ff.db.Exec(query, key, currentValue)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	// insert the old value
	for i := 0; i < version; i++ {
		query = `INSERT INTO oldValues (key, version, value) VALUES (?, ?, ?);`
		value, err := record.GetValue(i)
		if err != nil {
			return err
		}
		_, err = ff.db.Exec(query, key, i, value)
		if err != nil {
			return err
		}
	}

	ff.logger.Printf("Inserted record into SQLite Database Driver with key: %v", key)
	return nil
}

// insertMetadata - insert metadata into the database
func (ff *SQLiteDatabaseDriver) insertMetadata(record KvRecord) error {
	key := record.Key
	metadata := record.Metadata.GetAll()
	ff.logger.Printf("Inserting metadata into SQLite Database Driver with key: %v", key)
	// upsert the metadata
	for k, v := range metadata {
		query := `INSERT OR REPLACE INTO metadata (key, metadataKey, metadataValue) VALUES (?, ?, ?);`
		_, err := ff.db.Exec(query, key, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// findRecord - find a record in the database
func (ff *SQLiteDatabaseDriver) findRecord(key string) bool {
	query := `SELECT key FROM records WHERE key = ?;`
	rows, err := ff.db.Query(query, key)
	if err != nil {
		return false
	}
	defer rows.Close()
	return rows.Next()
}

// getRecord - get a record from the database
func (ff *SQLiteDatabaseDriver) getRecord(key string) (KvRecord, error) {
	query := `SELECT key, value FROM records WHERE key = ?;`
	rows, err := ff.db.Query(query, key)
	if err != nil {
		return KvRecord{}, err
	}
	defer rows.Close()
	if rows.Next() {
		var record KvRecord
		err := rows.Scan(&record.Key, &record.Value)
		if err != nil {
			return KvRecord{}, err
		}
		return record, nil
	}
	return KvRecord{}, nil
}

// getMetadata - get metadata from the database
func (ff *SQLiteDatabaseDriver) getMetadata(key string) (map[string]string, error) {
	query := `SELECT metadataKey, metadataValue FROM metadata WHERE key = ?;`
	rows, err := ff.db.Query(query, key)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	metadata := map[string]string{}
	for rows.Next() {
		var k, v string
		err := rows.Scan(&k, &v)
		if err != nil {
			return nil, err
		}
		metadata[k] = v
	}
	return metadata, nil
}

// deleteRecord - delete a record from the database
func (ff *SQLiteDatabaseDriver) deleteRecord(key string) error {
	// confirm ff.db is not nil
	if ff.db == nil {
		// open the database
		ff.initDatabase()
	}

	query := `DELETE FROM records WHERE key = ?;`
	_, err := ff.db.Exec(query, key)
	if err != nil {
		return err
	}
	return nil
}

// deleteMetadata - delete metadata from the database
func (ff *SQLiteDatabaseDriver) deleteMetadata(key string) error {
	// confirm ff.db is not nil
	if ff.db == nil {
		// open the database
		ff.initDatabase()
	}

	query := `DELETE FROM metadata WHERE key = ?;`
	_, err := ff.db.Exec(query, key)
	if err != nil {
		return err
	}
	return nil
}
