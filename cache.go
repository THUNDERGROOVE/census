package census

import (
	"time"
)

type Cache struct {
	invalid time.Time
}

func NewCache() Cache {
	return Cache{invalid: time.Now().Add(time.Minute * 2)}
}

func (c *Cache) IsInvalid() bool {
	if time.Now().After(c.invalid) {
		return true
	}
	return false
}
