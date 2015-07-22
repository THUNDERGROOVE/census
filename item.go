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
