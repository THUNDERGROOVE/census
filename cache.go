package census

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	ErrCacheNotFound = fmt.Errorf("No Cache was found")
)

// Cache is a struct
type Cache struct {
	invalid time.Time // Legacy @TODO: Remove when possible
	expires time.Time
}

// NewCache is going to change to NewCacheUpdate soon
func NewCache() Cache {
	return Cache{invalid: time.Now().Add(time.Minute * 2)}
}

// NewCacheUpdate soon to be NewCache returns a new Cache object that
func NewCacheUpdate(dur time.Duration) Cache {
	return Cache{
		expires: time.Now().Add(dur),
	}
}

// IsInvalid returns if the data needs to be requested again OR updated

func (c *Cache) IsInvalid() bool {
	if time.Now().After(c.expires) {
		return true
	}
	return false
}

// InvalidateIn invalidates the cache in the duration provided
func (c *Cache) InvalidateIn(dur time.Duration) {
	c.expires = time.Now().Add(dur)
}

// WriteCache writes the given interface to our caching filesystem
//
// @TODO: Maybe use a cacheType type with constants to aid in pulling cache?
// @TODO: Switch to encoding/gob for performance?
func WriteCache(cacheType string, identifier string, v interface{}) error {
	filename, path := cacheNames(cacheType, identifier)
	if err := os.MkdirAll(path, 0777); err == os.ErrPermission {
		return err
	}

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filename, data, 0777); err != nil {
		return err
	}
	return nil
}

// ReadCache reads the cache of the given type for the identifier and writes it into the interface
func ReadCache(cacheType, identifier string, v interface{}) error {
	filename, _ := cacheNames(cacheType, identifier)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

func cacheNames(cacheType, identifier string) (filename string, path string) {
	filename = fmt.Sprintf("cache/%v/%v", cacheType, identifier)
	path = fmt.Sprintf("cache/%v/", cacheType)
	return filename, path
}
