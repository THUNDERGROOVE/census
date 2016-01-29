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
	"testing"
	"time"
)

type testCache struct {
	Cache `json:"cache"`

	Name string `json:"name"`
	ID   int    `json:"id"`
}

func TestRedisCache(t *testing.T) {

	tc1 := new(testCache)
	tc2 := new(testCache)

	tc1.Cache = NewCacheUpdate(time.Hour)

	tc1.Name = "John"
	tc1.ID = 1

	
	c := new(Census)
	c.cacheType = CensusCacheRedis

	c.redisURL = "redis://127.0.0.1:6379"

	if err := c.RedisConnect(); err != nil {
		t.Errorf("RedisConnect: %s", err.Error())
		t.Fail()
	}

	
	if err := c.WriteCache(CACHE_TEST, tc1.ID, tc1); err != nil {
		t.Fatalf("Failed to write Cache: %s\n", err.Error())
	}

	if err := c.ReadCache(CACHE_TEST, tc1.ID, tc2); err != nil {
		t.Fatalf("Failed to read Cache: %s\n", err.Error())
	}

	if tc1.ID != tc2.ID || tc1.Name != tc2.Name {
		t.Error(tc2)
		t.Fatalf("Cache didn't match?")
	}
}
