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

type Population struct {
	VS uint
	TR uint
	NC uint
}

func (p *Population) Total() int {
	return int(p.VS + p.TR + p.NC)
}

func (p *Population) VSPercent() int {
	return int(float64(p.VS) / float64(p.Total()) * 100)
}

func (p *Population) TRPercent() int {
	return int(float64(p.TR) / float64(p.Total()) * 100)
}
func (p *Population) NCPercent() int {
	return int(float64(p.NC) / float64(p.Total()) * 100)
}

type PopulationSet struct {
	Servers map[string]*Population
	parent  *Census
}

func (c *Census) NewPopulationSet() *PopulationSet {
	pop := &PopulationSet{
		Servers: make(map[string]*Population),
		parent:  c,
	}
	return pop
}
func (c *PopulationSet) PlayerLogin(server, factionID string) {
	if _, ok := c.Servers[server]; !ok {
		c.Servers[server] = new(Population)
	}
	switch factionID {
	case VS:
		c.Servers[server].VS += 1
	case TR:
		c.Servers[server].TR += 1
	case NC:
		c.Servers[server].NC += 1
	default:
		fmt.Printf("PlayerLogin called with factionID that wasn't known: %v\n", factionID)
	}
}

func (c *PopulationSet) PlayerLogout(server, factionID string) {
	if _, ok := c.Servers[server]; !ok {
		c.Servers[server] = new(Population)
	}
	switch factionID {
	case VS:
		c.Servers[server].VS -= 1
	case TR:
		c.Servers[server].TR -= 1
	case NC:
		c.Servers[server].NC -= 1
	default:
		fmt.Printf("PlayerLogout called with factionID that wasn't known: %v\n", factionID)
	}
}
