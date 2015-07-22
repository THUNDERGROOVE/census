package census

import (
	"fmt"
	"testing"
)

func TestGetName(t *testing.T) {
	c := NewCensus("s:maximumtwang", "ps2ps4us:v2")
	fmt.Println(c.GetItemName(70966, "en"))
}
