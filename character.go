package census

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var ErrCharDoesNotExist = fmt.Errorf("census: That character doesn't exist")

// Character is a struct that contains all available information for a character in Planetside 2
// We've factored out all of the date constants and instead convert it using the unix timestamps given.
// it leads to quicker conversions
type Characters struct {
	Characters []Character `json:"character_list"`
}

// Character is a struct representing a character in the Census API with all possible useful resolves
type Character struct {
	CensusData
	Cache
	ID string `json:"character_id"`

	Name struct {
		First string `json:"first"`
		Lower string `json:"lower"`
	} `json:"name"`

	FactionID string `json:"faction_id"`

	TitleID string `json:"title_id"`

	Times struct {
		Creation      string `json:"creation"`
		LastSave      string `json:"last_save"`
		LastLogin     string `json:"last_login"`
		LoginCount    string `json:"login_count"`
		MinutesPlayed string `json:"minutes_played"`
	} `json:"times"`

	Certs struct {
		Earned        string `json:"earned_points"`
		Gifted        string `json:"gifted_points"`
		Spent         string `json:"spent_points"`
		Available     string `json:"available_points"`
		PercentToNext string `json:"percent_to_next"`
	} `json:"certs"`

	Battlerank struct {
		Rank          string `json:"value"`
		PercentToNext string `json:"percent_to_next"`
	} `json:"battle_rank"`

	DailyRibbon struct {
		Count string `json:"count"`
		Time  string `json:"time"`
	} `json:"daily_ribbon"`

	ProfileID string `json:"profile_id"`

	Outfit struct {
		ID          string `json:"outfit_id"`
		Name        string `json:"name"`
		Alias       string `json:"alias"`
		LeaderID    string `json:"leader_character_id"`
		MemberCount string `json:"member_count"`
		TimeCreated string `json:"time_created"`
	} `json:"outfit"`

	OnlineStatus string `json:"online_status"`

	Stats struct {
		Stat []struct {
			Name            string `json:"stat_name"`
			ProfileID       string `json:"profile_id"`
			ValueForever    string `json:"value_forever"`
			ValueMonthly    string `json:"value_monthly"`
			ValueWeekly     string `json:"value_weekly"`
			ValueDaily      string `json:"value_daily"`
			ValueOneLifeMax string `json:"value_one_life_max"`
			LastSave        string `json:"last_save"`
		} `json:"stat"`
		StatHistory []struct {
			Name       string            `json:"stat_name"`
			AllTime    string            `json:"all_time"`
			OneLifeMax string            `json:"one_life_max"`
			Day        map[string]string `json:"day"`
			Month      map[string]string `json:"Month"`
			Week       map[string]string `json:"Week"`
		} `json:"stat_history"`
		WeaponStat []struct {
			Name      string `json:"weapon_deaths"`
			ItemID    string `json:"item_id"`
			VehicleID string `json:"vehicle_id"`
			Value     string `json:"value"`
			LastSave  string `json:"last_save"`
		} `json:"weapon_stat"`
	} `json:"stats"`
	Faction struct {
		Name struct {
			En string `json:"en"`
			De string `json:"de"`
			Es string `json:"es"`
			Fr string `json:"fr"`
			It string `json:"it"`
			Tr string `json:"tr"`
		} `json:"name"`
		ImageSetID     string `json:"image_set_id"`
		ImageID        string `json:"image_id"`
		ImagePath      string `json:"image_path"`
		CodeTag        string `json"code_tag"`
		UserSelectable string `json:"user_selectable"`
	} `json:"faction"`

	Items []struct {
		ID         string `json:"item_id"`
		StackCount string `json:"stack_count"`
	} `json:"items"`

	FriendsList []struct {
		ID            string `json:"character_id"`
		LastLoginTime string `json:"last_login_time"`
		Onlint        string `json:"online"`
	} `json:"friends_list"`

	World string `json:"world_id"`

	Parent   *Census `json:"-"`
	IsCached bool    `json:"-"`
}

// GetCharacterByName is a method to get a character instance by name
// it internally caches information and will pull new information if
// fifteen minutes has passed since the update.
func (c *Census) GetCharacterByName(name string) (*Character, error) {
	name = strings.ToLower(name)
	char := new(Character)
	id, err := c.GetCharacterID(name)
	if err != nil {
		return nil, err
	}
	// @BUG: logic error?
	if err := ReadCache(CACHE_CHARACTER, id, char); err != ErrCacheNotFound {
		if !char.Cache.IsInvalid() {
			return char, nil
		}
	}

	return c.getChar(name)
}

const character_resolves = "item,profile,faction,stat,weapon_stat,stat_history,online_status,friends,world,outfit"

func (c *Census) getChar(name string) (*Character, error) {
	req := c.NewRequest(REQUEST_CHARACTER, "name.first_lower="+name,
		character_resolves,
		1)
	tmp := new(Characters)

	if err := req.Do(tmp); err != nil {
		return nil, err
	}

	if len(tmp.Characters) < 1 {
		return nil, ErrCharDoesNotExist
	}
	char := tmp.Characters[0]
	char.Parent = c
	char.Cache = NewCacheUpdate(time.Minute * 15)

	return &char, nil

}

func (c *Census) GetCharacterByID(ID string) (*Character, error) {
	chars := new(Characters)
	char := new(Character)
	if err := ReadCache(CACHE_CHARACTER, ID, char); err != ErrCacheNotFound {
		if !char.Cache.IsInvalid() {
			return char, nil
		}
	}

	req := c.NewRequest(REQUEST_CHARACTER, "character_id="+ID, character_resolves, 1)
	if err := req.Do(chars); err != nil {
		return nil, err
	}
	if len(chars.Characters) < 1 {
		return nil, ErrCharDoesNotExist
	}
	char = &chars.Characters[0]
	char.Parent = c
	char.Cache = NewCacheUpdate(time.Minute * 15)
	return char, nil
}

// GetCharacterID
// @TODO: Update to use Cache if possible
func (c *Census) GetCharacterID(name string) (string, error) {
	chars := new(Characters)

	req := c.NewRequest(
		REQUEST_CHARACTER,
		"name.first_lower="+strings.ToLower(name), "", 1, "c:show=name")

	if err := req.Do(chars); err != nil {
		return "<nil>", err
	}

	if len(chars.Characters) == 0 {
		return "<nil>", ErrCharDoesNotExist
	}

	char := chars.Characters[0]
	return char.Name.First, nil
}

// GetFacilitiesCaptured returns the facilities captured via the stats history
func (c *Character) GetFacilitiesCaptured() int {
	return c.getStatFromStatHistory("facility_capture")
}

func (c *Character) GetScore() int {
	return c.getStatFromStatHistory("score")
}

func (c *Character) GetMedals() int {
	return c.getStatFromStatHistory("medals")
}

func (c *Character) GetRibbons() int {
	return c.getStatFromStatHistory("ribbons")
}

func (c *Character) GetCerts() int {
	i, _ := strconv.Atoi(c.Certs.Gifted)
	return c.getStatFromStatHistory("certs") + i
}

func (c *Character) GetFacilitiesDefended() int {
	return c.getStatFromStatHistory("facility_defend")
}

// GetKills returns the total kill count for a given Character
func (c *Character) GetKills() int {
	return c.getStatFromStatHistory("kills")
}

// GetDeaths returns the total death count for a given Character
func (c *Character) GetDeaths() int {
	return c.getStatFromStatHistory("deaths")
}

// KDR returns the KDR of a Character
func (c *Character) KDR() float64 {
	return float64(c.GetKills()) / float64(c.GetDeaths())
}

// KDRS returns the KDR of a Character in a more human
// readable format
func (c *Character) KDRS() string {
	return strconv.FormatFloat(c.KDR(), 'f', 2, 64)
}

// ServerName returns the English name for the server the
// Character resides on
func (c *Character) ServerName() string {
	s := c.Parent.GetServerByID(c.World)
	return s.Name.En
}

// TKPercent is the percent of TKs in the last 1000 kills.
// This is changing to total once my caching system is in place.
// @TODO: Error handle? Nahh
func (c *Character) TKPercent() int {
	events, err := c.Parent.GetAllKillEvents(c.ID)
	if err != nil { // Log maybe?
		return -100
	}
	var tkcount int
	for _, v := range events.List {
		if v.Character.FactionID == c.FactionID {
			tkcount += 1
		}
	}
	// tk/k * 100
	return int(
		float64(tkcount) / float64(len(events.List)) * 100,
	)
}

/* Helpers
 */

func (c *Character) getStatFromStatHistory(s string) int {
	for _, stat := range c.Stats.StatHistory {
		if stat.Name == s {
			i, _ := strconv.Atoi(stat.AllTime)
			return i
		}
	}
	return 0
}
