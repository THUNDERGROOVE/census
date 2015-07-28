package census

import (
	"fmt"
	"testing"
	"time"
)

func TestEventStreaming(t *testing.T) {
	c := NewCensus("s:maximumtwang", "ps2ps4us:v2")
	events := c.NewEventStream()
	time.Sleep(3 * time.Second)
	if err := events.ClearSubscriptions(); err != nil {
		fmt.Printf("err: %v\n", err.Error())
	}
	sub := NewEventSubscription()
	sub.Worlds = []string{"all"}
	sub.Characters = []string{"all"}
	sub.EventNames = []string{"PlayerLogin", "PlayerLogout"}
	if err := events.Subscribe(sub); err != nil {
		t.Fatalf("failed to subscribe: %v\n", err.Error())
	}
	for {
		select {
		case err := <-events.Err:
			fmt.Printf("error in websocket world: %v\n", err.Error())
		case <-events.Closed:
			fmt.Printf("Connection closed\n")
			break
		case ev := <-events.Events:
			fmt.Printf("Event: %v\n", ev)
		}
	}
}
