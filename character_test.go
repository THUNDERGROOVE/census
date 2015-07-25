package census

import (
	"fmt"
	"testing"
)

func TestQueryCharacter(t *testing.T) {
	c := NewCensus("s:maximumtwang", "ps2ps4us:v2")
	char, err := c.GetCharacterByName("THUNDERGROOVE")
	if err != nil {
		t.Fatalf("Error getting character information: %v\n", err.Error())
	}

	fmt.Printf("ID: %v\n", char.ID)
	for _, v := range char.Stats.StatHistory {
		fmt.Printf("Stat: [%v]: [%v]\n", v.Name, v.AllTime)
	}

	fmt.Printf("[%v] %v@%v\n", char.Outfit.Alias, char.Name.First, char.ServerName())
}

func TestKillCount(t *testing.T) {
	c := NewCensus("s:maximumtwang", "ps2ps4us:v2")
	char, err := c.GetCharacterByName("THUNDERGROOVE")
	if err != nil {
		t.Fatalf("Error getting character information: %v\n", err.Error())
	}
	fmt.Printf("Got %v kills\n", char.KDR())
}
