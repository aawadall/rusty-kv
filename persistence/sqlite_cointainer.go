package persistence

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

// SqliteContainer - sqlite container
type SqliteContainer struct {
	dbLocation string

	mu sync.Mutex
}

// NewSqliteContainer - create a new sqlite container
func NewSqliteContainer(dbLocation string) *SqliteContainer {

	return &SqliteContainer{
		dbLocation: dbLocation,
		mu:         sync.Mutex{},
	}
}

// Unified Execute Query Function
func (sc *SqliteContainer) ExecuteQuery(query string, args ...interface{}) (*sql.Rows, error) {
	// Lock the mutex
	sc.mu.Lock()
	defer sc.mu.Unlock()

	// inspect args
	for _, arg := range args {
		log.Printf("arg: %v (%T)", arg, arg)
	}
	// open the database
	db, err := sql.Open("sqlite3", sc.dbLocation)

	// if there is an error, return it
	if err != nil {
		return nil, err
	}

	// defer closing the database
	defer db.Close()

	// execute the query
	rows, err := db.Query(query, args...)

	// if there is an error, return it
	if err != nil {
		return nil, err
	}

	// return the rows
	return rows, nil

}

// Initialize Database
func (sc *SqliteContainer) InitDatabase(tables []string) error {
	// Lock the mutex
	sc.mu.Lock()
	defer sc.mu.Unlock()

	// open the database
	db, err := sql.Open("sqlite3", sc.dbLocation)

	// if there is an error, return it
	if err != nil {
		return err
	}

	// defer closing the database
	defer db.Close()

	// execute the query
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			return err
		}
	}

	// return nil
	return nil
}

// *sql.Rows rows affected
func RowsAffected(sr *sql.Rows) (int64, error) {
	// TODO: implement this
	return 0, nil
}
