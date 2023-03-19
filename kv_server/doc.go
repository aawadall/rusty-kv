package kvserver

/*
	Package kvserver - This is the KVServer package for the application.
	it is responsible for keeping track of the key-value pairs
	it is also responsible for coordinating with the following packages
	- api - for handling the API requests
	- cluster - to coordinate with other servers
	- persistence - for frequent book keeping
	- config - for configuration management

	key value pairs will have the following struct shape:
	- key
	- value
	- metadata
		- version
		- last modified time
		- TTL
		- tenant

	key values will reside both in memory and on disk (persistence)
*/
