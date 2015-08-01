package census

import (
	"fmt"
)

type requestType string

const (
	REQUEST_CHARACTER        requestType = "character"
	REQUEST_CHARACTER_EVENTS requestType = "characters_event"
	REQUEST_WORLD            requestType = "world"
)

type Request struct {
	*Census
	url string
}

func (c *Census) NewRequest(Type requestType, query string, resolves string, limit int, more ...string) *Request {
	req := new(Request)
	req.Census = c

	base := fmt.Sprintf("%v%v/get/%v/%v/",
		BaseURL,
		c.serviceID,
		c.namespace, Type)

	if query != "" {
		base = fmt.Sprintf("%v?%v", base, query)
	}

	if resolves != "" {
		base = fmt.Sprintf("%v&c:resolve=%v", base, resolves)
	}
	if limit != 0 {
		base = fmt.Sprintf("%v&c:limit=%v", base, limit)
	}

	for _, v := range more {
		base = fmt.Sprintf("%v&%v", base, v)
	}

	req.url = base
	//	fmt.Printf("url: %v\n", base)
	return req
}

func (r *Request) Do(v interface{}) error {
	return decode(r.Census, r.url, v)
}
