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
