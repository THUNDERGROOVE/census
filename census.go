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

// The census package is used to query data from the Daybreak Game's census API
// for use in getting data for Planetside 2
package census

import (
	"strings"

	"github.com/garyburd/redigo/redis"
)

var BaseURLOld = "http://census.soe.com/"

type CensusCacheType int

const (
	CensusCacheNone CensusCacheType = iota
	CensusCacheDisk
	CensusCacheRedis
)

// CensusData is a struct that contains various metadata that a Census request can have.
type CensusData struct {
	Error string `json:"error"`
}

// Census is the main object you use to query data
type Census struct {
	serviceID string
	namespace string
	cacheType CensusCacheType
	redisURL  string
	conn      redis.Conn
}

// NewCensus returns a new census object given your service ID
func NewCensus(ServiceID string, Namespace string) *Census {
	c := new(Census)
	c.serviceID = ServiceID
	c.namespace = Namespace
	return c
}

// NewCensusCache allows you to specify how to cache.
func NewCensusCache(ServiceID string, Namespace string, cacheType CensusCacheType) *Census {
	c := new(Census)
	c.serviceID = ServiceID
	c.namespace = Namespace
	c.cacheType = cacheType
	return c
}

// CleanNamespace returns a proper namespace for internal use in queries
func (c *Census) CleanNamespace() string {
	if strings.Contains(c.namespace, ":") {
		return strings.Split(c.namespace, ":")[0]
	}
	return c.namespace
}

func (c *Census) IsEU() bool {
	if strings.Contains(c.namespace, "eu") {
		return true
	}
	return false
}
