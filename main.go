package main

import (
	kvserver "github.com/aawadall/simple-kv/kv_server"
)

func main() {
	// TODO - Entry point code  here
	// define server
	server := kvserver.NewKVServer()

	server.Start()
}
