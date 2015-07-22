package census

import (
	"fmt"
	"testing"
)

func TestQueryCharacter(t *testing.T) {
	c := NewCensus("s:maximumtwang", "ps2ps4us:v2")
	char, err := c.QueryCharacterByExactName("THUNDERGROOVE")
	if err != nil {
		t.Fatalf("Error getting character information: %v\n", err.Error())
	}
	fmt.Printf("[%v] %v\n", char.Outfit.Alias, char.Name.First)
}

func TestKillCount(t *testing.T) {
	c := NewCensus("s:maximumtwang", "ps2ps4us:v2")
	char, err := c.QueryCharacterByExactName("THUNDERGROOVE")
	if err != nil {
		t.Fatalf("Error getting character information: %v\n", err.Error())
	}
	fmt.Printf("Got %v kills\n", char.KDR())
}
