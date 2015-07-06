package census

import (
	"fmt"
	"testing"
)

func TestGetCollection(t *testing.T) {
	c := NewCensus("s:maximumtwang", "ps2ps4us:v2")
	cols, err := getCollection(c)
	if err != nil {
		t.Fatalf("Error getting collections: %v\n", err.Error())
	}
	for _, v := range cols {
		fmt.Printf("%#v\n", v)
	}
}
