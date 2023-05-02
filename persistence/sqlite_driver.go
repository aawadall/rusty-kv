package persistence

import (
	"database/sql"
	"log"
	"os"

	"github.com/aawadall/simple-kv/types"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// init SQL statements
var sqlInit = []string{
	`CREATE TABLE IF NOT EXISTS records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT UNIQUE,
		value BLOB
	);`,
	`CREATE TABLE IF NOT EXISTS metadata (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT,
		metadataKey TEXT,
		metadataValue TEXT,
		FOREIGN KEY(key) REFERENCES records(key),
		UNIQUE (key, metadataKey)
	);`,
	`CREATE TABLE IF NOT EXISTS oldValues (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT,
		version INTEGER,
		value BLOB,
		FOREIGN KEY(key) REFERENCES records(key),
		UNIQUE (key, version)
	);`,
}

var sqlOperations = map[string]string{
	"insertRecord":      `INSERT OR REPLACE INTO records (key, value) VALUES (?, ?);`,
	"insertOldValue":    `INSERT INTO oldValues (key, version, value) VALUES (?, ?, ?);`,
	"insertMetadata":    `INSERT OR REPLACE INTO metadata (key, metadataKey, metadataValue) VALUES (?, ?, ?);`,
	"selectRecordByKey": `SELECT key FROM records WHERE key = ?;`,
	"selectAllRecords":  `SELECT key FROM records;`,
}

// helper struct
type Transaction struct {
	TokenId string
	Query   string
	Args    []interface{}
}

// SQLite Driver
type SQLiteDriver struct {
	dbLocation string
	logger     *log.Logger
}

// NewSQLiteDriver - create a new sqlite driver
func NewSQLiteDriver(dbLocation string) *SQLiteDriver {
	driver := &SQLiteDriver{
		dbLocation: dbLocation,
		logger:     log.New(os.Stdout, "sqlite: ", log.LstdFlags),
	}

	driver.logger.Printf("Creating SQLite Driver with location: %v", dbLocation)

	// initialize the driver
	driver.init()

	return driver
}

// init - initialize the driver
func (driver *SQLiteDriver) init() {
	// open the database
	db, err := sql.Open("sqlite3", driver.dbLocation)

	if err != nil {
		driver.logger.Printf("Error opening database: %v", err.Error())
	}

	defer db.Close()

	// initialize the database
	for _, query := range sqlInit {
		_, err := db.Exec(query)

		if err != nil {
			driver.logger.Printf("Error initializing database: %v", err.Error())
		}
	}
}

// Implement the Driver interface
// Write - write a record to the database
func (driver *SQLiteDriver) Write(record *KvRecord) error {
	// mementos
	transactions := []Transaction{}

	// insert record
	transaction, err := driver.insertRecord(record)

	// append the token to the transaction

	transactions = append(transactions, transaction)
	if err != nil {
		driver.logger.Printf("Error inserting record: %v", err.Error())
		driver.rollback(transactions)
		return err
	}

	// insert old values
	old_value_transactions, err := driver.insertOldValues(record)

	// append old value transactions to the transactions
	transactions = append(transactions, old_value_transactions...)

	if err != nil {
		driver.logger.Printf("Error inserting old values: %v", err.Error())
		driver.rollback(transactions)
		return err
	}

	// insert metadata
	metadata_transaction, err := driver.insertMetadata(record)
	transactions = append(transactions, metadata_transaction...)
	if err != nil {
		driver.logger.Printf("Error inserting metadata: %v", err.Error())
		driver.rollback(transactions)
		return err
	}

	return nil
}

// Read - read a record from the database
func (driver *SQLiteDriver) Read(key string) (*KvRecord, error) {
	record := &KvRecord{
		Key: key,
	}

	// get the record
	err := driver.getRecord(record)
	if err != nil {
		driver.logger.Printf("Error getting record: %v", err.Error())
		return nil, err
	}

	// get the old values
	err = driver.getOldValues(record)
	if err != nil {
		driver.logger.Printf("Error getting old values: %v", err.Error())
		return nil, err
	}

	// get the metadata
	err = driver.getMetadata(record)
	if err != nil {
		driver.logger.Printf("Error getting metadata: %v", err.Error())
		return nil, err
	}

	return record, nil
}

// Delete - delete a record from the database
func (driver *SQLiteDriver) Delete(key string) error {
	// open the database
	db, err := sql.Open("sqlite3", driver.dbLocation)

	if err != nil {
		driver.logger.Printf("Error opening database: %v", err.Error())
		return err
	}

	defer db.Close()

	// delete the record
	_, err = db.Exec(sqlOperations["deleteRecord"], key)
	if err != nil {
		driver.logger.Printf("Error deleting record: %v", err.Error())
		return err
	}

	return nil
}

// Compare - compare a record to the database
func (driver *SQLiteDriver) Compare(record *KvRecord) (bool, error) {
	// get the record
	dbRecord, err := driver.Read(record.Key)
	if err != nil {
		driver.logger.Printf("Error getting record: %v", err.Error())
		return false, err
	}

	// compare the records
	return matchRecords(record, dbRecord), nil
}

// Load - load all records from the database
func (driver *SQLiteDriver) Load() ([]*KvRecord, error) {
	// open the database
	db, err := sql.Open("sqlite3", driver.dbLocation)

	if err != nil {
		driver.logger.Printf("Error opening database: %v", err.Error())
		return nil, err
	}

	defer db.Close()

	// get all records
	rows, err := db.Query(sqlOperations["selectAllRecords"])
	if err != nil {
		driver.logger.Printf("Error selecting records: %v", err.Error())
		return nil, err
	}

	defer rows.Close()

	// load the records
	records := []*KvRecord{}
	for rows.Next() {
		var key string
		err := rows.Scan(&key)
		if err != nil {
			driver.logger.Printf("Error scanning record: %v", err.Error())
			return nil, err
		}

		record, err := driver.Read(key)
		if err != nil {
			driver.logger.Printf("Error reading record: %v", err.Error())
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

// helper functions
// insertRecord - insert a record into the database
func (driver *SQLiteDriver) insertRecord(record *KvRecord) (Transaction, error) {
	value, err := record.Value.Get(-1)
	if err != nil {
		driver.logger.Printf("Error getting value: %v", err.Error())
		return Transaction{}, err
	}

	return driver.insert(sqlOperations["insertRecord"], record.Key, value)
}

// insertOldValues - insert old values into the database
func (driver *SQLiteDriver) insertOldValues(record *KvRecord) ([]Transaction, error) {
	// get version
	version := record.Value.GetVersion()
	transactions := []Transaction{}
	// insert old values
	for i := 0; i < version; i++ {
		value, err := record.Value.Get(i)
		if err != nil {
			driver.logger.Printf("Error getting value: %v", err.Error())
			return []Transaction{}, err
		}
		token, err := driver.insert(sqlOperations["insertOldValue"], record.Key, value)
		if err != nil {
			driver.logger.Printf("Error inserting old value: %v", err.Error())
			return []Transaction{}, err
		}
		transactions = append(transactions, token)
	}

	return transactions, nil
}

// insertMetadata - insert metadata into the database
func (driver *SQLiteDriver) insertMetadata(record *KvRecord) ([]Transaction, error) {
	transactions := []Transaction{}

	// insert metadata
	for key, value := range record.Metadata.GetAll() {
		byte_value := []byte(value)
		transaction, err := driver.insert(sqlOperations["insertMetadata"], key, byte_value)
		if err != nil {
			driver.logger.Printf("Error inserting metadata: %v", err.Error())
			return []Transaction{}, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

// insert - insert a record into the database
func (driver *SQLiteDriver) insert(query string, key string, value []byte) (Transaction, error) {
	// open the database
	db, err := sql.Open("sqlite3", driver.dbLocation)

	if err != nil {
		driver.logger.Printf("Error opening database: %v", err.Error())
		return Transaction{}, err
	}

	defer db.Close()

	// insert the record
	result, err := db.Exec(query, key, value)
	if err != nil {
		driver.logger.Printf("Error inserting record: %v", err.Error())
		return Transaction{}, err
	}

	// get the token
	_, err = result.LastInsertId()
	if err != nil {
		driver.logger.Printf("Error getting token: %v", err.Error())
		return Transaction{}, err
	}

	token := makeToken()

	transaction := Transaction{
		TokenId: token,
		Query:   query,
		Args:    []interface{}{key, value},
	}
	return transaction, nil
}

// rollback - rollback a transaction
func (driver *SQLiteDriver) rollback(transactions []Transaction) {
	// treat transaction as a stack
	for len(transactions) > 0 {
		// pop the transaction
		transaction := transactions[len(transactions)-1]

		// reverse the token
		driver.reverseTransaction(transaction)

		// remove the token
		transactions = transactions[:len(transactions)-1]
	}
}

// get record
func (driver *SQLiteDriver) getRecord(key *KvRecord) (error) {
	// open the database
	db, err := sql.Open("sqlite3", driver.dbLocation)

	if err != nil {
		driver.logger.Printf("Error opening database: %v", err.Error())
		return err
	}

	defer db.Close()

	// get the record
	rows, err := db.Query(sqlOperations["selectRecord"], key)
	if err != nil {
		driver.logger.Printf("Error selecting record: %v", err.Error())
		return err
	}

	defer rows.Close()

	// load the record
	record := KvRecord{}
	for rows.Next() {
		var key string
		var value []byte
		err := rows.Scan(&key, &value)
		if err != nil {
			driver.logger.Printf("Error scanning record: %v", err.Error())
			return err
		}

		record.Key = key
		record.Value = &types.ValuesContainer{}
		record.Value.Set(value)
	}

	return nil
}


// get old values 
func (driver *SQLiteDriver) getOldValues(key *KvRecord) (error) {
	// open the database
	db, err := sql.Open("sqlite3", driver.dbLocation)

	if err != nil {
		driver.logger.Printf("Error opening database: %v", err.Error())
		return err
	}

	defer db.Close()

	// get the record
	rows, err := db.Query(sqlOperations["selectOldValues"], key)
	if err != nil {
		driver.logger.Printf("Error selecting old values: %v", err.Error())
		return err
	}

	defer rows.Close()

	// load the record
	for rows.Next() {
		var key string
		var value []byte
		err := rows.Scan(&key, &value)
		if err != nil {
			driver.logger.Printf("Error scanning old values: %v", err.Error())
			return err
		}

		record := KvRecord{}
		record.Key = key
		record.Value = &types.ValuesContainer{}
		record.Value.Set(value)
	}

	return nil
}
// helper functions
// make token
func makeToken() string {
	// UUID
	uuid := uuid.New()
	return uuid.String()
}
