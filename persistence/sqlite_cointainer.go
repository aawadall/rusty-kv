package persistence

import (
	"database/sql"
	"sync"
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

// *sql.Rows rows affected
func RowsAffected(sr *sql.Rows) (int64, error) {
	// TODO: implement this
	return 0, nil
}
