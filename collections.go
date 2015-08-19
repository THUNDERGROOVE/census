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

// Collection represents a single collection
type Collection struct {
	Count        int
	Dynamic      bool
	UnknownCount bool
	Hidden       bool
	Name         string
	ResolveList  []string
}

func GetCollections(c *Census) ([]*Collection, error) {
	var out []*Collection
	tmp := map[string]interface{}{}

	url := BaseURL + "get/" + c.namespace
	fmt.Printf("Getting url: %v\n", url)
	if err := decode(c, url, &tmp); err != nil {
		return nil, err
	}
	jq := jsonq.NewQuery(tmp)
	n, err := jq.ArrayOfObjects("datatype_list")
	if err != nil {
		return nil, err
	}

	//fmt.Printf("%v\n", data)
	for _, v := range n {
		data := jsonq.NewQuery(v)
		out = append(out, parse_collection(data))
	}
	return out, nil
}

func parse_collection(jq *jsonq.JsonQuery) *Collection {
	out := new(Collection)
	out.Count, _ = jq.Int("count")
	out.Hidden, _ = jq.Bool("hidden")
	out.Name, _ = jq.String("name")
	out.ResolveList, _ = jq.ArrayOfStrings("resolve_list")
	return out
}
