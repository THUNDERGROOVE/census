package census

import (
	"testing"
	"time"
)

type testCache struct {
	Cache `json:"cache":`

	Name string `json:"name"`
	ID   int    `json:"id"`
}

func TestCacheWrite(t *testing.T) {
	c := new(testCache)
	c2 := new(testCache)

	c.ID = 1337
	c.Name = "L33TH4x0r69"
	c.Cache = NewCacheUpdate(time.Minute * 30)

	if err := WriteCache(CACHE_TEST, c.ID, c); err != nil {
		t.Fatal("Failed to write cache to disk: [%v]", err.Error())
	}

	if err := ReadCache(CACHE_TEST, c.ID, c2); err != nil {
		t.Fatalf("Failed to read cache from disk: [%v]", err.Error())
	}

	if c.ID != c2.ID || c.Name != c2.Name {
		t.Fatalf("Cache didn't match?")
	}
}
