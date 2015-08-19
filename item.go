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
	"github.com/jmoiron/jsonq"
)

// GetItemName returns the name of a given id in the language provided
func (c *Census) GetItemName(id int, lang string) string {
	tmp := map[string]interface{}{}
	url := fmt.Sprintf("%vget/%v/item/?item_id=%v", BaseURL, c.namespace, id)
	if err := decode(c, url, &tmp); err != nil {
		return err.Error()
	}
	jq := jsonq.NewQuery(tmp)
	a, _ := jq.ArrayOfObjects("item_list")
	item := a[0]
	q := jsonq.NewQuery(item)
	if lang == "" {
		lang = "en"
	}
	s, _ := q.String("name", lang)
	return s
}
