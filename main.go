package main

import (
	"fmt"

	kvserver "github.com/aawadall/simple-kv/kv_server"
	"github.com/aawadall/simple-kv/persistence"
	"github.com/aawadall/simple-kv/types"
)

func main() {
	// Entry point code  here

	// Define Features Manager
	featureFlagManager := NewFeatureFlagManager()

	featureFlagManager.Add("direct_sqlite", "Use the sqlite driver directly", true)
	featureFlagManager.Add("kv_server", "Instantiate and use KV Server", false)

	// Experiment: directly use sqlite driver
	if featureFlagManager.IsEnabled("direct_sqlite") {
		location := "local.sqlite"
		driver := persistence.NewSQLiteDatabaseDriver(location)

		// prepare record
		record := types.NewKVRecord("test", []byte("test"))

		// inspect record
		fmt.Println("Inspecting Record: ")
		fmt.Printf("Record ID: %s\n", record.Id)
		fmt.Printf("Record Key: %s\n", record.Key)

		value, err := record.Value.Get(-1)
		fmt.Printf("Record Value: %v\n", value)

		// write a record
		err = driver.Write(*record)

		if err != nil {
			fmt.Println(err)
		}

		// read a record
		readRecord, err := driver.Read("test")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(readRecord)

		// delete a record
		err = driver.Delete("test")

		if err != nil {
			fmt.Println(err)
		}

	}
	// define server
	if featureFlagManager.IsEnabled("kv_server") {
		config := map[string]string{
			"driver":        "sqlite",
			"file_location": "kv.sqlite",
			"sync_interval": "120",
		}
		server := kvserver.NewKVServer(config)

		server.Start()
	}
}
