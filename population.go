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
