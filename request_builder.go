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
	"fmt"
)

type requestType string

const (
	REQUEST_CHARACTER        requestType = "character"
	REQUEST_CHARACTER_EVENTS requestType = "characters_event"
	REQUEST_WORLD            requestType = "world"
)

type Request struct {
	*Census
	url string
}

func (c *Census) NewRequest(Type requestType, query string, resolves string, limit int, more ...string) *Request {
	req := new(Request)
	req.Census = c

	base := fmt.Sprintf("%v%v/get/%v/%v/",
		BaseURL,
		c.serviceID,
		c.namespace, Type)

	if query != "" {
		base = fmt.Sprintf("%v?%v", base, query)
	}

	if resolves != "" {
		base = fmt.Sprintf("%v&c:resolve=%v", base, resolves)
	}
	if limit != 0 {
		base = fmt.Sprintf("%v&c:limit=%v", base, limit)
	}

	for _, v := range more {
		base = fmt.Sprintf("%v&%v", base, v)
	}

	req.url = base
	//	fmt.Printf("url: %v\n", base)
	return req
}

func (r *Request) Do(v interface{}) error {
	return decode(r.Census, r.url, v)
}
