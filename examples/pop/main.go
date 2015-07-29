// pop shows how to use the streaming API and PopulationSet struct
// to poll for population data
package main

import (
	"bufio"
	"fmt"
	"github.com/THUNDERGROOVE/census"
	"os"
)

func main() {
	fmt.Printf("Starting to gather server population info\nPress Ctrl-C to quit\n")
	c := census.NewCensus("s:maximumtwang", "ps2ps4us:v2")

	events := c.NewEventStream()
	sub := census.NewEventSubscription()
	sub.Worlds = []string{"all"}
	sub.Characters = []string{"all"}
	sub.EventNames = []string{"PlayerLogin", "PlayerLogout"}
	if err := events.Subscribe(sub); err != nil {
		fmt.Printf("FAIL: Couldn't subscribe to events: [%v]\n", err.Error())
		return
	}
	pop := c.NewPopulationSet()

	infoChan := make(chan struct{}, 0)
	go func() {
		fmt.Printf("Press Enter to print stats!\n")
		for {
			pause()
			infoChan <- struct{}{}
		}
	}()

	for {
		select {
		case <-infoChan:
			fmt.Printf("Printing server population info: \n")
			for name, server := range pop.Servers {
				fmt.Printf("%v:\n", name)
				fmt.Printf("VS: %v:%%%v\n", server.VS, server.VSPercent())
				fmt.Printf("TR: %v:%%%v\n", server.TR, server.TRPercent())
				fmt.Printf("NC: %v:%%%v\n", server.NC, server.NCPercent())
			}
		case err := <-events.Err:
			fmt.Printf("error: %v\n", err.Error())
		case <-events.Closed:
			fmt.Printf("Websocket closed\n")
			break
		case event := <-events.Events:
			switch event.Payload.EventName {
			case "PlayerLogin":
				ch, err := c.GetCharacterByID(event.Payload.CharacterID)
				if err != nil {
					fmt.Printf("ERROR: Failed to get character from ID: '%v' [%v]\n",
						event.Payload.CharacterID, err.Error())
					continue
				}
				server := c.GetServerByID(event.Payload.WorldID)
				pop.PlayerLogin(server.Name.En, ch.FactionID)
			case "PlayerLogout":
				ch, err := c.GetCharacterByID(event.Payload.CharacterID)
				if err != nil {
					fmt.Printf("ERROR: Failed to get character from ID: '%v' [%v]\n",
						event.Payload.CharacterID, err.Error())
					continue
				}
				server := c.GetServerByID(event.Payload.WorldID)
				pop.PlayerLogin(server.Name.En, ch.FactionID)
			}
		}
	}
}

func pause() {
	bufio.NewReader(os.Stdin).ReadString('\n')
}
