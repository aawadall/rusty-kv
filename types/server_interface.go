package types

// Server API interface
type Server interface {
	GetStatus() (interface{}, error)
}
