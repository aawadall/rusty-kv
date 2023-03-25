package types

// Server API interface
type Server interface {
	GetStatus() (interface{}, error)
	Start()
	Stop()
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Delete(key string) error
	SetMetadata(key string, metadataKey string, metadataValue string) error
	GetMetadata(key string, metadataKey string) (string, error)
	DeleteMetadata(key string, metadataKey string) error
	GetAllMetadata(key string) (map[string]string, error)
	Find(partialKey string) ([]string, error)
	FindByMetadata(query string) ([]string, error)
}
