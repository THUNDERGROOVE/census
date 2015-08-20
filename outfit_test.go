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
	"testing"
)

const _testOutfit = `{"outfit_list":[{"outfit_id":"37531502290403331","name":"The Abyss","name_lower":"the abyss","alias":"ABYS","alias_lower":"abys","time_created":"1437401448","time_created_date":"2015-07-20 14:10:48.0","leader_character_id":"5428352933368439697","member_count":"244"}],"returned":1}`

func TestDecodeOutfit(t *testing.T) {
	outfits := new(Outfits)
	if err := decode([]byte(_testOutfit), outfits); err != nil {
		t.Fatalf("failed on good JSON: %v", err.Error())
	}
	if len(outfits.Outfits) != 1 {
		t.Fatal("should have only gotten one outfit")
	}
	outfit := outfits.Outfits[0]
	if outfit.NameLower != "the abyss" {
		t.Fatalf("unexpected name: %v", outfit.NameLower)
	}
}
