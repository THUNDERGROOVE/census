package census

import (
	"encoding/json"
	"fmt"
	"github.com/lafikl/fluent"
	"io/ioutil"
)

var ErrTooManyRetries = fmt.Errorf("Failed too many retries")

// This looks ugly as fuck
func decode(c *Census, url string, v interface{}) error {
	t := fluent.New()

	t.Get(url).Retry(3)
	if resp, err := t.Send(); err == nil {
		defer resp.Body.Close()
		if data, err := ioutil.ReadAll(resp.Body); err == nil {
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}
