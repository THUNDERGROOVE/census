package census

import (
	"fmt"
	"log"
	"time"
)

const (
	VS = "1"
	NC = "2"
	TR = "3"
)

// CharacterEvent is a single frame of a /characters_event/ request
type CharacterEvent struct {
	Character          Character `json:"character"`
	CharacterID        string    `json:"character_id"`
	AttackerID         string    `json:"attacker_character_id"`
	IsHeadshot         string    `json:"is_headshot"`
	IsCrit             string    `json:"is_critical"`
	AttackerWeaponID   string    `json:"attacker_weapon_id"`
	AttackerVehicleID  string    `json:"attacker_vehicle_id"`
	Time               string    `json:"timestamp"`
	Zone               string    `json:"zone_id"`
	World              string    `json:"world_id"`
	CharacterLoadoutID string    `json:"character_loadout_id"`
	AttackerLoadoutID  string    `json:"attacker_loadout_id"`
	TableType          string    `json:"table_type"`
}

// CharacterEventList is a struct capable of being Unmarshaled from a
// /characters_event/ request
type CharacterEventList struct {
	Cache    `json:"cache"`
	List     []CharacterEvent `json:"characters_event_list"`
	Returned int              `json:"returned"`
}

// GetKillEvents returns a CharacterEventList given a count and characterID
func (c *Census) GetKillEvents(count int, characterID string) *CharacterEventList {
	out := new(CharacterEventList)

	req := c.NewRequest(
		REQUEST_CHARACTER_EVENTS,
		"character_id="+characterID,
		"character",
		count,
		"type=KILL")

	if err := req.Do(out); err != nil {
		log.Printf("ERROR: GetKillEvents() -> req.Do() [%v]", err.Error())
		return nil
	}

	return out
}

func (c *Census) GetAllKillEvents(characterID string) (*CharacterEventList, error) {
	out := new(CharacterEventList)

	if err := ReadCache(CACHE_CHARACTER_EVENTS, "kills"+characterID, out); err == nil {

	}

	// Perform operatons

	if err := WriteCache(CACHE_CHARACTER_EVENTS, "kills"+characterID, out); err != nil {
		return out, err
	}

	return out, nil
}

func (c *CharacterEventList) UpdateCache(census *Census, characterID string) error {
	if !c.Cache.IsInvalid() {
		return nil
	}
	// Request the last 1000 kills since the last pull
	req := census.NewRequest(REQUEST_CHARACTER_EVENTS,
		"character_id="+characterID,
		"character",
		1000,
		"type=KILL",
		fmt.Sprintf("after=%v",
			c.Cache.LastUpdated.Unix()))

	if err := req.Do(c); err != nil {
		return err
	}

	if c.Returned == 1000 {
		// @TODO: Make run a complete rebuild of the kills cache
		panic("Been too long since we updated.  Over 1000; cache invalid.  Must rebuild")
		return nil
	}

	c.Cache = NewCacheUpdate(time.Minute * 30)
	c.LastUpdated = time.Now()
	return nil
}

func (c *Character) TeamKillsInLast(count int) int {
	events := c.Parent.GetKillEvents(count, c.ID)
	var kills int
	for _, v := range events.List {
		if v.Character.FactionID == c.FactionID {
			kills += 1
		}
	}
	return kills
}
