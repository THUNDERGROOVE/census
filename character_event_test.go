package census

import (
	"fmt"
	"testing"
)

func TestGetKillEvents(t *testing.T) {
	fmt.Printf("Creating new census instance\n")
	c := NewCensus("s:maximumtwang", "ps2ps4us:v2")
	char, err := c.GetCharacterByName("THUNDERGROOVE")
	fmt.Printf("Getting character to find ID\n")
	if err != nil {
		t.Fatalf("Error getting character information: %v\n", err.Error())
	}
	fmt.Printf("Getting 10 kill events\n")
	events := c.GetKillEvents(10, char.ID)
	fmt.Printf("Got %v events\n", len(events.List))
	for _, event := range events.List {
		fmt.Printf("Killed: %v@%v\n", event.Character.Name.First, event.Character.FactionID)
	}

	fmt.Printf("%v has teamkilled %v in the last 100 kills", char.Name.First, char.TeamKillsInLast(100))
}
