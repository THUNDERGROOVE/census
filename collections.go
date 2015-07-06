package census

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// json_collection is used to circumvent the inconsistancies in the census
// data
type json_collection struct {
	count        json.RawMessage `json:"count"`
	hidden       json.RawMessage `json:"hidden"`
	name         string          `json:"name"`
	resolve_list []string        `json:"resolve_list"`
}

type json_collection_parse struct {
	datatypelist []json_collection `json:"datatype_list"`
}

// Collection represents a single collection
type Collection struct {
	Count        int
	Dynamic      bool
	UnknownCount bool
	Hidden       bool
	Name         string
	ResolveList  []string
}

func (j *json_collection) ToCollection() *Collection {
	out := new(Collection)

	switch string(j.count) {
	case "?":
		out.UnknownCount = true
	case "dynamic":
		out.Dynamic = true
	default:
		i, err := strconv.Atoi(string(j.count))
		if err != nil {
			fmt.Println("Unexpected library error:",
				" switch hit default but not convertable to integer",
				err.Error())
		}
		out.Count = i
	}

	out.Hidden, _ = strconv.ParseBool(string(j.hidden))

	out.Name = string(j.name)
	out.ResolveList = j.resolve_list

	return out
}

func getCollection(c *Census) ([]*Collection, error) {
	var out []*Collection
	tmp := new(json_collection_parse)
	tmp.datatypelist = make([]json_collection, 0)
	url := BaseURL + "get/" + c.namespace
	fmt.Printf("Getting url: %v\n", url)
	if err := decode(c, url, tmp); err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", tmp)
	for _, v := range tmp.datatypelist {
		out = append(out, v.ToCollection())
	}
	return out, nil
}
