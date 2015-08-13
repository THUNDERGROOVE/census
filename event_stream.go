package census

import (
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
	"golang.org/x/net/websocket"
)

type eventType uint8

// Cosntants not used..  May do some stuff with these in the future
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

// EventStream is an abstraction for the Planetside 2 streaming data API
// instead of using a raw websocket connection you can just use channels.
type EventStream struct {
	parent *Census
	conn   *websocket.Conn
	Err    chan error
	Events chan Event
	Closed chan struct{}
}

// GlobalDecoder exists so we don't allocate a new decoder every response
var GlobalDecoder = ffjson.NewDecoder()

// Event is an event from the Planetside real time event streaming API.
// @TODO: Add the rest of the fields.  They're all in the documentation
type Event struct {
	Payload EventPayload
	Service string `json:"service"`
	Type    string `json:"type"`
}

type EventPayload struct {
	EventName   string `json:"event_name"`
	Time        string `json:"timestamp"`
	CharacterID string `json:"character_id"`
	WorldID     string `json:"world_id"`
}

// EventSent is a representation of any data we send to the API
type EventSent struct {
	Service    string   `json:"service"`
	Action     string   `json:"action"`
	Worlds     []string `json:"worlds"`
	Characters []string `json:"characters"`
	EventNames []string `json:"eventNames"`
	//	All        string   `json:"all"`
}

// NewEventSubscription returns an EventSent with the required fields for an
// event subscription already filled out
func NewEventSubscription() *EventSent {
	s := new(EventSent)
	s.Service = "event"
	s.Action = "subscribe"
	s.Characters = []string{}
	s.Worlds = []string{}
	s.EventNames = []string{}
	return s
}

// NewEventStream returns an EventStream
//
// NOTICE: This method dials a websocket
//       : This method starts a Go routine
func (c *Census) NewEventStream() *EventStream {
	ev := new(EventStream)
	ev.parent = c
	ev.Events = make(chan Event, 0)
	ev.Err = make(chan error, 0)
	ev.Closed = make(chan struct{}, 0)

	var err error
	url := fmt.Sprintf("wss://push.planetside2.com/streaming?environment=%v&service-id=%v", c.CleanNamespace(), c.serviceID)
	ev.conn, err = websocket.Dial(
		url,
		"", "http://localhost/")
	if err != nil {
		ev.Err <- err
		ev.Closed <- struct{}{}
		return ev
	}

	go func() {
		// Only keep one allocation around.  Pass values to the channel instead.
		var event = new(Event)
		var buf = make([]byte, 2048)
		for {
			/*
				if err := websocket.JSON.Receive(ev.conn, event); err != nil {
					ev.Err <- err
					ev.Closed <- struct{}{}
					break
				}
				ev.Events <- *event*/

			if n, err := ev.conn.Read(buf); err != nil {
				ev.Err <- err
				ev.Closed <- struct{}{}
				break
			} else {
				data := buf[:n]
				if err := GlobalDecoder.DecodeFast(data, event); err != nil {
					ev.Err <- err
					ev.Closed <- struct{}{}
					break
				}

				ev.Events <- *event
			}
		}
	}()
	return ev
}

// Subscribe verifies that the provided EventSent is a subscripton event
// and sends it
//
// @TODO: Need checks to make sure it's a subscribe
func (e *EventStream) Subscribe(sub *EventSent) error {
	return e.RawEventSent(sub)
}

// RawEventSent sends a Raw, user formed EventSent
func (e *EventStream) RawEventSent(sent *EventSent) error {
	return websocket.JSON.Send(e.conn, sent)

}

// ClearSubscriptions sends an event to clear all event subscriptions
func (e *EventStream) ClearSubscriptions() error {
	return e.RawEventSent(&EventSent{
		Service: "event",
		Action:  "clearSubscribe",
		/*All:     "true"*/})
}

// Close closes the unlderlying websocket.
// This will send a struct{}{} down the Closed channel
func (c *EventStream) Close() {
	c.conn.Close()
}
