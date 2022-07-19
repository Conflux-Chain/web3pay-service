package store

import (
	"errors"
	"io"
	"time"
)

var (
	// custom store errors
	ErrNotFound = errors.New("not found")
)

// Store is implemented by any object that persist data.
type Store interface {
	io.Closer

	// Check if record not in store by error
	IsRecordNotFound(err error) bool
}

type CacheStore interface {
	Store

	// IncrBy increase an item of type int64 by n. Returns an error if the item's value is
	// not an int64, or if it was not found. If there is no error, the incremented
	// value is returned.
	IncrBy(k string, n int64) (int64, error)
	// Get gets an item from the cache. Returns the item or nil, and a bool indicating
	// whether the key was found.
	Get(k string) (interface{}, bool)
	// Set adds an item to the cache, replacing any existing item. If the duration is 0
	// (DefaultExpiration), the cache's default expiration time is used. If it is -1
	// (NoExpiration), the item never expires.
	Set(k string, x interface{}, d time.Duration)
	// Delete an item from the cache. Does nothing if the key is not in the cache.
	Delete(k string)
	// Flush deletes all items from the cache.
	Flush() error
}
