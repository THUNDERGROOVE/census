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

func TestCacheWrite(t *testing.T) {
	c := new(testCache)
	c2 := new(testCache)

	c.ID = 1337
	c.Name = "L33TH4x0r69"
	c.Cache = NewCacheUpdate(time.Minute * 30)

	if err := WriteCache(CACHE_TEST, c.ID, c); err != nil {
		t.Fatalf("Failed to write cache to disk: [%v]", err.Error())
	}

	if err := ReadCache(CACHE_TEST, c.ID, c2); err != nil {
		t.Fatalf("Failed to read cache from disk: [%v]", err.Error())
	}

	if c.ID != c2.ID || c.Name != c2.Name {
		t.Fatalf("Cache didn't match?")
	}
}
