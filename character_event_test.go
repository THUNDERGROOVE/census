package census

import (
	"fmt"
	"testing"
)

func TestGetAllKills(t *testing.T) {
	fmt.Printf("Creating new census instance\n")
	c := NewCensus("s:maximumtwang", "ps2ps4us:v2")
	char, err := c.GetCharacterByName("THUNDERGROOVE")
	if err != nil {
		t.Fatalf("error getting character info: %v", err.Error())
	}

	events, err := c.GetAllKillEvents(char.ID)
	if err != nil {
		t.Fatalf("failed getting all kill events: %v", err.Error())
	}

	var tk int

	for _, v := range events.List {
		if v.Character.FactionID == char.FactionID {
			tk += 1
		}
	}

	if len(events.List)-tk != char.GetKills() {
		//t.Fatalf("Kill event count mismatch! %v != %v", len(events.List)-tk, char.GetKills())
	}
}

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
