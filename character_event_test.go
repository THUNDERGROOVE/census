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
