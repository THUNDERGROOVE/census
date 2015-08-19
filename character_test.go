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
