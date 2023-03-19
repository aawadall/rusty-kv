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

	// sleep for a bit to simulate real stopping

	time.Sleep(15 * time.Second)

	server.Stop()

	// wait for goroutines to finish
	wg.Wait()
}
