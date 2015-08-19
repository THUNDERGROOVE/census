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

package census

import (
	"fmt"
)

var ErrOutfitNotExist = fmt.Errorf("census: the outfit doesn't exist")

type Outfit struct {
	ID           string `json:"outfit_id"`
	Name         string `json:"name"`
	NameLower    string `json:"name_lower"`
	Alias        string `json:"alias"`
	AliasLower   string `json:"alias_lower"`
	TimeCreated  string `json:"time_created"`
	LeaderCharID string `json:"leader_character_id"`
	MemberCount  string `json:"member_count"`
}

type Outfits struct {
	CensusData
	Outfits []Outfit `json:"outfit_list"`
}

// GetOutfitByName returns an outfit instance if it exists, otherwise it returns an error
func (c *Census) GetOutfitByName(name string) (*Outfit, error) {
	req := c.NewRequest(REQUEST_OUTFIT, "name="+name, "", 1)
	outfits := &Outfits{Outfits: []Outfit{}}
	if err := req.Do(outfits); err != nil {
		return nil, err
	}

	if len(outfits.Outfits) == 0 {
		return nil, ErrOutfitNotExist
	}

	outfit := &outfits.Outfits[0]
	return outfit, nil
}
