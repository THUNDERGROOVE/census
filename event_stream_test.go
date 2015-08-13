package census

import (
	"github.com/pquerna/ffjson/ffjson"
	"testing"
)

func BenchmarkUnmarshal(b *testing.B) {
	testData := []byte(`{"payload":{"character_id":"5428366097940635073","event_name":"PlayerLogout","timestamp":"1439505006","world_id":"1002"},"service":"event","type":"serviceMessage"}`)

	dec := ffjson.NewDecoder()
	var event Event
	for i := 0; i < 100000; i++ {
		b.Log("Decoding", i)
		if err := dec.Decode(testData, &event); err != nil {
			b.Fatalf("decode error: %v\n", err.Error())
		}
	}
}

/*
func TestEventStreaming(t *testing.T) {

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	c := NewCensus("s:maximumtwang", "ps2ps4us:v2")
	events := c.NewEventStream()
	defer events.Close()
	time.Sleep(3 * time.Second)
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
			fmt.Printf("\nerror in websocket world: %v\n", err.Error())
		case <-events.Closed:
			fmt.Printf("Connection closed\n")
			break
		case ev := <-events.Events:
			switch ev.Payload.EventName {
			case "PlayerLogin":
				fallthrough
			case "PlayerLogout":
				ch, err := c.GetCharacterByID(ev.Payload.CharacterID)
				if err != nil {
					fmt.Printf("ERROR: Couldn't get character by ID [%v]\n", err.Error())
				}
				server := c.GetServerByID(ev.Payload.WorldID)
				fmt.Printf("%v: %v %v\n", ev.Payload.EventName, ch.Name.First, server.Name.En)
			}
		}
	}
}*/
