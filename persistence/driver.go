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
