package main

import (
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
	expected := "value1"
	server.Set("key1", []byte(expected))
	val, err := server.Get("key1")

	if err != nil {
		panic(err)
	}

	// cast value to string
	sVal := string(val.([]byte))

	// print both
	println("Expected: ", expected)
	println("Actual: ", sVal)

	// sleep for a bit to simulate real stopping

	time.Sleep(15 * time.Second)

	server.Stop()

	// wait for goroutines to finish
	wg.Wait()
}
