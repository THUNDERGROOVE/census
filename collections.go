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
