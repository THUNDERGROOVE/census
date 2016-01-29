// The MIT License (MIT)
//
// Copyright (c) 2015 Nick Powell
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
//
//

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
	ErrNoCache       = fmt.Errorf("census: Cache method called with cache disabled")
)

type cacheType uint8

const (
	CACHE_CHARACTER cacheType = iota
	CACHE_CHARACTER_EVENTS

	CACHE_TEST
)

var CacheTypes = []string{
	"character",
	"character_events",
	"tests",
}


// TODO: Going to need to break up the data stored in cache into two parts
// We need one part being the Cache struct and an additional for the actual data.
// Otherwise, the performance improvements from using redis will likly be negligable
// considering most time will be spent decoding JSON anyways in the edge cases
// where we're just checking if the cache is valid 

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
func (c *Census) WriteCache(ct cacheType, id interface{}, v interface{}) error {
	switch c.cacheType {
	case CensusCacheNone:
		return ErrNoCache
	case CensusCacheDisk:
		return c.writeCacheDisk(ct, id, v)
	case CensusCacheRedis:
		return c.writeCacheRedis(ct, id, v)
	default:
		panic("UNKNOWN CACHE TYPE")
	}
}

func (c *Census) writeCacheDisk(ct cacheType, id interface{}, v interface{}) error {
	filename, path := cacheNames(ct, id)
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

func (c *Census) writeCacheRedis(ct cacheType, id interface{}, v interface{}) error {
	filename, _ := cacheNames(ct, id)
	// Maybe want a better setup for keys

	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	
	_, err = c.conn.Do("SET", filename, data)
	if err != nil {
		return err
	}
	return nil
}

// ReadCache reads the cache of the given type for the identifier and writes
// into the given interface
func (c *Census) ReadCache(ct cacheType, id interface{}, v interface{}) error {
	switch c.cacheType {
	case CensusCacheNone:
		return ErrNoCache
	case CensusCacheDisk:
		return c.readCacheDisk(ct, id, v)
	case CensusCacheRedis:
		return c.readCacheRedis(ct, id, v)
	default:
		panic("UNKNOWN CACHE TYPE HIT ReadCache")
	}
}

func (c *Census) readCacheDisk(ct cacheType, id interface{}, v interface{}) error {
	filename, _ := cacheNames(ct, id)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

func (c *Census) readCacheRedis(ct cacheType, id interface{}, v interface{}) error {
	filename, _ := cacheNames(ct, id)
	data, err := c.conn.Do("GET", filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(data.([]uint8)), v); err != nil {
		return err
	}
	return nil
}

/* No longer using.  Possibly will reimplement it
func CheckCache(ct cacheType, identifier interface{}) bool {
	filename, _ := cacheNames(ct, identifier)

	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}*/

// cacheNames is a helper function to provide a filename and path given a
// cacheType and an idenitfier
func cacheNames(ct cacheType, identifier interface{}) (filename string, path string) {
	filename = fmt.Sprintf("cache/%s/%v", CacheTypes[ct], identifier)
	path = fmt.Sprintf("cache/%s/", CacheTypes[ct])
	return filename, path
}
