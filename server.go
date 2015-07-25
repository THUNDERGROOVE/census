package census

import (
	"log"
)

type Server struct {
	ID    string `json:"world_id"`
	State string `json:"state"`
	Name  struct {
		En string `json:"en"`
	} `json:"name"`
}

type Servers struct {
	Servers []Server `json:"world_list"`
}

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
