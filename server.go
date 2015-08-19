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
	"log"
)

// Server represents a single server or in other contexts a world
type Server struct {
	ID    string `json:"world_id"`
	State string `json:"state"`
	Name  struct {
		En string `json:"en"`
	} `json:"name"`
}

// Servers is a group of servers.  Usually would be used to parse a server from
// a census response
type Servers struct {
	CensusData
	Servers []Server `json:"world_list"`
}

// GetServerByID returns a server by a given ID.
//
// TODO: This may need breaking changes to indicated an error
func (c *Census) GetServerByID(id string) Server {
	req := c.NewRequest(REQUEST_WORLD, "world_id="+id, "", 0)
	s := new(Servers)
	if err := req.Do(s); err != nil {
		log.Printf("Error decoding servers: [%v]", err.Error())
	}
	if len(s.Servers) >= 1 {
		return s.Servers[0]
	} else {
		log.Printf("No server returned")
		return Server{}
	}
}
