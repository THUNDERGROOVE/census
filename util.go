package census

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// This looks ugly as fuck
func decode(c *Census, url string, v *json_collection_parse) error {
	if resp, err := http.Get(url); err == nil {
		if data, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Printf("Got response: %v with size %v\n", resp.Status, len(data))
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
			fmt.Printf("%v\n", v)
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}
