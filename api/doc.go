package api

/*
	Package api - This is the API package for the application.
	it is responsible for handling the API requests, in the following protocols:
	- REST over HTTP/HTTPS
	- GRPC

*/

/*
	Proposed API:
	- get value for a key
	- set value for a key
	- delete a key
	- get all keys with optional
		- prefix
		- limit
		- metadata
			- TTL comparison
			- last modified time comparison
			- Version comparison
	- check if a key exists
	- system health check
	- update configuration
	- get configuration
	- get cluster information
	- get cluster health
	- get cluster configuration
	- get cluster members
*/
