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
	"testing"
)

const _testChar = `{"character_list":[{"character_id":"5428352933374094753","name":{"first":"THUNDERGROOVE","first_lower":"thundergroove"},"faction_id":"1","head_id":"1","title_id":"0","times":{"creation":"1435089505","creation_date":"2015-06-23 19:58:25.0","last_save":"1439968580","last_save_date":"2015-08-19 07:16:20.0","last_login":"1439958152","last_login_date":"2015-08-19 04:22:32.0","login_count":"283","minutes_played":"19640"},"certs":{"earned_points":"29762","gifted_points":"4016","spent_points":"33143","available_points":"635","percent_to_next":"0.9920000000007"},"battle_rank":{"percent_to_next":"33","value":"73"},"profile_id":"17","daily_ribbon":{"count":"0","time":"1439967600","date":"2015-08-19 07:00:00.0"}}],"returned":1}`

// static decoding
func TestDecodeCharacter(t *testing.T) {
	chars := new(Characters)

	if err := decode([]byte(_testChar), chars); err != nil {
		t.Fatalf("failed to decode good JSON: %v\n", err.Error())
	}
	if len(chars.Characters) != 1 {
		t.Fatal("incorrect amount of characters parsed")
	}
	char := chars.Characters[0]

	if char.Name.Lower != "thundergroove" {
		t.Fatalf("failed to get value from json got: '%v'", char.Name.Lower)
	}
}

// dynamic tests.  Requres census to be working

func TestGetCharacterByName(t *testing.T) {
	_, err := testingCensus.GetCharacterByName("THUNDERGROOVE")
	if err != nil {
		t.Fatalf("couldn't find THUNDERGROOVE: %v", err.Error())
	}
}

func TestGetCharacterByNameFail(t *testing.T) {
	_, err := testingCensus.GetCharacterByName("names can't contain spaces")
	if err == nil {
		t.Fatalf("expected error.  Didn't get it")
	}
}

func TestGetChar(t *testing.T) {
	_, err := testingCensus.getChar("thundergroove") //works on lower
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetCharFail(t *testing.T) {
	_, err := testingCensus.getChar("names can't contain spaces")
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGetCharacterByID(t *testing.T) {
	_, err := testingCensus.GetCharacterByID("5428352933374094753")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetCharacterID(t *testing.T) {
	id, err := testingCensus.GetCharacterID("THUNDERGROOVE")
	if err != nil {
		t.Fatal(err.Error())
	}
	if id != "5428352933374094753" {
		t.Logf("%v != 5428352933374094753", id)
		t.Fatal("ID mismatch")
	}
}
