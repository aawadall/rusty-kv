package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// experiments related to main project above this folder

type User struct {
	Name  string
	Email string
}

func main() {
	// define sqlite3 database connection
	dbLocation := "exp.sqlite"

	// define a new table
	CreateTable(dbLocation)

	// insert a new row
	user := User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	InsertRow(dbLocation, user)
	// select all rows

	// select a row

	// update a row

	// delete a row
}

// CreateTable creates a new table
func CreateTable(dbLocation string) {
	// open database connection
	db, err := sql.Open("sqlite3", dbLocation)

	// check for errors
	if err != nil {
		log.Fatal(err)
	}

	// close database connection
	defer db.Close()

	// define query
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	)`

	// execute query
	result, err := db.Exec(query)

	// check for errors
	if err != nil {
		log.Fatal(err)
	}

	// inspect result
	log.Println("Result Stats")
	rowsAffected, _ := result.RowsAffected()
	log.Printf("Rows Affected: %d", rowsAffected)
	lastInsertID, _ := result.LastInsertId()
	log.Printf("Last Insert ID: %d", lastInsertID)
}

// InsertRow inserts a new row
func InsertRow(dbLocation string, user User) {
	// open database connection
	db, err := sql.Open("sqlite3", dbLocation)

	// check for errors
	if err != nil {
		log.Fatal(err)
	}

	// close database connection
	defer db.Close()

	// define query
	query := `INSERT INTO users (name, email) VALUES (?, ?)`

	// define arguments
	name := user.Name
	email := user.Email

	// execute query
	result, err := db.Exec(query, name, email)

	// check for errors
	if err != nil {
		log.Fatal(err)
	}

	// inspect result
	log.Println("Result Stats")
	rowsAffected, _ := result.RowsAffected()
	log.Printf("Rows Affected: %d", rowsAffected)
	lastInsertID, _ := result.LastInsertId()
	log.Printf("Last Insert ID: %d", lastInsertID)
}
