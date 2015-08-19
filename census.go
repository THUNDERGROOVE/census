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

// The census package is used to query data from the census API.
//
// It's centered more so around data from Planetside 2
package census

import (
	"strings"
)

var BaseURLOld = "http://census.soe.com/"

// CensusData is a struct that contains various metadata that a Census request can have.
type CensusData struct {
	Error string `json:"error"`
}

// NewCensus returns a new census object given your service ID
func NewCensus(ServiceID string, Namespace string) *Census {
	c := new(Census)
	c.serviceID = ServiceID
	c.namespace = Namespace
	return c
}

// Census is the main object you use to query data
type Census struct {
	serviceID string
	namespace string
}

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
