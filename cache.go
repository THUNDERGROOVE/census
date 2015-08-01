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

type cacheType uint8

const (
	CACHE_CHARACTER cacheType = iota
	CACHE_CHARACTER_EVENTS

	CACHE_TEST
)

// Cache is a struct
type Cache struct {
	invalid     time.Time // Legacy @TODO: Remove when possible
	Expires     time.Time `json:"expires"`
	LastUpdated time.Time `json:"last-updated"`
}

// NewCache is going to change to NewCacheUpdate soon
func NewCache() Cache {
	return Cache{invalid: time.Now().Add(time.Minute * 2)}
}

// NewCacheUpdate soon to be NewCache returns a new Cache object that
func NewCacheUpdate(dur time.Duration) Cache {
	return Cache{
		Expires: time.Now().Add(dur),
	}
}

// IsInvalid returns if the data needs to be requested again OR updated
func (c *Cache) IsInvalid() bool {
	if time.Now().After(c.Expires) {
		return true
	}
	return false
}

// InvalidateIn invalidates the cache in the duration provided
func (c *Cache) InvalidateIn(dur time.Duration) {
	c.Expires = time.Now().Add(dur)
}

// WriteCache writes the given interface to our caching filesystem
//
// @TODO: Maybe use a cacheType type with constants to aid in pulling cache?
// @TODO: Switch to encoding/gob for performance?
//        I hear msgpack performance is good: github.com/vmihailenco/msgpack
//        gob has bad performance, but has very powerful reflection powers.
func WriteCache(cacheType cacheType, identifier interface{}, v interface{}) error {
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

// ReadCache reads the cache of the given type for the identifier and writes
// into the given interface
func ReadCache(ct cacheType, identifier interface{}, v interface{}) error {
	filename, _ := cacheNames(ct, identifier)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

func CheckCache(ct cacheType, identifier interface{}) bool {
	filename, _ := cacheNames(ct, identifier)

	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}

// cacheNames is a helper function to provide a filename and path given a
// cacheType and an idenitfier
func cacheNames(ct cacheType, identifier interface{}) (filename string, path string) {
	filename = fmt.Sprintf("cache/%v/%v", ct, identifier)
	path = fmt.Sprintf("cache/%v/", ct)
	return filename, path
}
