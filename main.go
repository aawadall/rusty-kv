package main

import (
	"fmt"
	"sync"
	"time"

	kvserver "github.com/aawadall/simple-kv/kv_server"
)

func main() {
	// TODO - Entry point code  here
	// define server
	server := kvserver.NewKVServer()

	// define wait group
	var wg sync.WaitGroup

	// start server using goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Start()
	}()

	// invoke a couple of operations
	// case 1 - create and get
	println("Case 1 - create and get")
	key := "key1"
	expected := "value1"
	server.Set(key, []byte(expected))
	val, err := server.Get(key)

	if err != nil {
		panic(err)
	}

	// cast value to string
	sVal := string(val.([]byte))

	// print both
	println("Expected: ", expected)
	println("Actual: ", sVal)

	// case 2 - delete and get
	println("Case 2 - delete and get")
	server.Delete(key)
	_, err = server.Get(key)

	println("we should get an error here")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())

	}

	// case 3 - create and update and get
	println("Case 3 - create and update and get")
	expected = "value2"
	key = "key2"
	server.Set(key, []byte("value1"))
	val, err = server.Get(key)

	if err != nil {
		panic(err)
	}

	sVal = string(val.([]byte))

	// print both
	println("before update")
	println("Actual: ", sVal)

	server.Set(key, []byte(expected))
	val, err = server.Get(key)

	if err != nil {
		panic(err)
	}

	sVal = string(val.([]byte))
	println("after update")
	println("Actual: ", sVal)
	println("Expected: ", expected)

	server.Delete(key)
	// case 4 - set metadata and get
	println("Case 4 - set metadata and get")
	key = "key3"
	expected = "value3"
	server.Set(key, []byte(expected))

	// set metadata
	server.SetMetadata(key, "metadata", "metadataValue")

	// get metadata
	metadata, err := server.GetMetadata(key, "metadata")

	if err != nil {
		panic(err)
	}

	println("Metadata: ", metadata)

	server.Delete(key)
	// sleep for a bit to simulate real stopping

	time.Sleep(15 * time.Second)

	server.Stop()

	// wait for goroutines to finish
	wg.Wait()
}
