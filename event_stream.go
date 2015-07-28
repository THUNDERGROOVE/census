package census

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
)

type eventType uint8

const (
	// Player Events
	EventNil eventType = iota
	EventAcheivementEarned
	EventBattleRankUp
	EventDeath
	EventItemAdded
	EventSkillAdded
	EventVehicleDestory
	EventGainXP
	EventFacilityCapture
	EventFacilityDefend

	// World Events
	EventContinentLock
	EventContinentUnlock
	EventFacilityControl
	EventAlertEvent

	// Both
	EventPlayerLogin
	EventPlayerLogout
)

type EventStream struct {
	parent *Census
	conn   *websocket.Conn
	Err    chan error
	Events chan Event
	Closed chan struct{}
}

type Event struct {
	Payload struct {
		EventName   string `json:"event_name"`
		Time        string `json:"timestamp"`
		CharacterID string `json:"character_id"`
		WorldID     string `json:"world_id"`
	} `json:"payload"`
	Service string `json:"service"`
	Type    string `json:"type"`
}

type EventSent struct {
	Service    string   `json:"service"`
	Action     string   `json:"action"`
	Worlds     []string `json:"worlds"`
	Characters []string `json:"characters"`
	EventNames []string `json:"eventNames"`
	All        string   `json:"all"`
}

func NewEventSubscription() *EventSent {
	s := new(EventSent)
	s.Service = "event"
	s.Action = "subscribe"
	return s
}

func (c *Census) NewEventStream() *EventStream {
	ev := new(EventStream)
	ev.parent = c
	ev.Err = make(chan error, 0)
	ev.Closed = make(chan struct{}, 0)

	var err error
	fmt.Printf("Starting websocket\n")
	url := fmt.Sprintf("wss://push.planetside2.com/streaming?environment=%v&service-id=%v", c.CleanNamespace(), c.serviceID)
	fmt.Printf("url: %v\n", url)
	ev.conn, err = websocket.Dial(
		url,
		"", "http://localhost/")
	if err != nil {
		ev.Err <- err
		ev.Closed <- struct{}{}
		return ev
	}

	fmt.Printf("Starting listen routine\n")
	go func() {

		var buffer = make([]byte, 2048)
		for {
			n, err := ev.conn.Read(buffer)
			if err != nil {
				ev.Err <- err
				continue
			}
			msg := buffer[:n]
			fmt.Printf("msg: %v\n", string(msg))
			event := Event{}
			if err := json.Unmarshal(msg, &event); err != nil {
				ev.Err <- err
				continue
			}
			ev.Events <- event
		}
	}()
	return ev
}

// @TODO: Need checks to make sure it's a subscribe
func (e *EventStream) Subscribe(sub *EventSent) error {
	return e.RawEventSent(sub)
}

func (e *EventStream) RawEventSent(sent *EventSent) error {
	data, err := json.Marshal(sent)
	if err != nil {
		return err
	}
	_, err = e.conn.Write(data)
	if err != nil {
		return err
	}
	return nil

}

func (e *EventStream) ClearSubscriptions() error {
	return e.RawEventSent(&EventSent{
		Service: "event",
		Action:  "clearSubscribe",
		All:     "true"})
}
