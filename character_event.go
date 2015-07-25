package census

import (
	"log"
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
	Cache              `json:"cache"`
}

// CharacterEventList is a struct capable of being Unmarshaled from a
// /characters_event/ request
type CharacterEventList struct {
	List []CharacterEvent `json:"characters_event_list"`
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
