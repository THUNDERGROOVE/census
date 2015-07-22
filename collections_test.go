package census

import (
	"fmt"
	"testing"
)

func TestGetCollection(t *testing.T) {
	c := NewCensus("s:maximumtwang", "ps2ps4us:v2")
	cols, err := GetCollections(c)
	if err != nil {
		t.Fatalf("Error getting collections: %v\n", err.Error())
	}
	fmt.Printf("Collections:\n")
	for _, v := range cols {
		fmt.Printf("%v\n", v.Name)
		for _, vv := range v.ResolveList {
			fmt.Printf("  %v\n", vv)
		}
	}
}
