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
	"encoding/json"
	"fmt"
	"github.com/lafikl/fluent"
	"io/ioutil"
)

var ErrTooManyRetries = fmt.Errorf("Failed too many retries")

// This looks ugly as fuck
func decodeURL(url string, v interface{}) error {
	t := fluent.New()

	t.Get(url).Retry(3)
	if resp, err := t.Send(); err == nil {
		defer resp.Body.Close()
		if data, err := ioutil.ReadAll(resp.Body); err == nil {
			return decode(data, v)
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}

func decode(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
