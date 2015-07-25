package census

import (
	"fmt"
	"log"
)

const (
	VS = "1"
	NC = "2"
	TR = "3"
)

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
	Cache `json:"cache"`
}

type CharacterEventList struct {
	List []CharacterEvent `json:"characters_event_list"`
}

func (c *Census) GetKillEvents(count int, characterID string) *CharacterEventList {
	out := new(CharacterEventList)

	url := fmt.Sprintf("%v%v/get/%v/characters_event/?character_id=%v&type=KILL&c:resolve=character&c:limit=%v", BaseURL, c.serviceID, c.namespace, characterID, count)
	fmt.Printf("url: %v\n", url)
	if err := decode(c, url, out); err != nil {
		log.Printf("error: %v", err.Error())
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
