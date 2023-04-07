package persistence

import "github.com/aawadall/simple-kv/types"

// Aliases
type KvRecord = types.KVRecord

// Persistence Driver
type Driver interface {
	Write(KvRecord) error
	Read(string) (KvRecord, error)
	Delete(string) error
	Compare(KvRecord) (bool, error)
	Load() ([]KvRecord, error)
}

// Helper Functions
func CompareRecords(a, b KvRecord) bool {
	return a.Key == b.Key && a.Value == b.Value && CompareMetadata(a.Metadata.GetAll(), b.Metadata.GetAll())
}

func CompareMetadata(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
