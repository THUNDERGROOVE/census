package census

import (
	"fmt"
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
	url := fmt.Sprintf("%v%v/get/%v/world/?world_id=%v", BaseURL, c.serviceID, c.namespace, id)
	fmt.Printf("URL: %v\n", url)
	s := new(Servers)
	if err := decode(c, url, s); err != nil {
		log.Printf("Error decoding servers: [%v]", err.Error())
	}

	if len(s.Servers) >= 1 {
		return s.Servers[0]
	} else {
		log.Printf("No server returned")
		return Server{}
	}
}
