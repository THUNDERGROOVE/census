package census

import (
	"fmt"
	"strconv"
	"strings"
)

var ErrCharDoesNotExist = fmt.Errorf("census: That character doesn't exist")

// Character is a struct that contains all available information for a character in Planetside 2
// We've factored out all of the date constants and instead convert it using the unix timestamps given.
// it leads to quicker conversions
type Characters struct {
	Characters []Character `json:"character_list"`
}

type Character struct {
	CensusData
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

var CharacterCahceMap = make(map[string]*CharacterCache)

type CharacterCache struct {
	Character
	Cache
}

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

func (c *Character) GetKills() int {
	return c.getStatFromStatHistory("kills")
}

func (c *Character) GetDeaths() int {
	return c.getStatFromStatHistory("deaths")
}

func (c *Character) KDR() float64 {
	return float64(c.GetKills()) / float64(c.GetDeaths())
}

func (c *Character) getStatFromStatHistory(s string) int {
	for _, stat := range c.Stats.StatHistory {
		if stat.Name == s {
			i, _ := strconv.Atoi(stat.AllTime)
			return i
		}
	}
	return 0
}

func (c *Character) ServerName() string {
	s := c.Parent.GetServerByID(c.World)
	return s.Name.En
}

func (c *Character) TKPercent() int {
	kills := c.TeamKillsInLast(1000)
	return int(float64(float64(kills)/1000) * 100)
}

func (c *Census) QueryCharacterByExactName(name string) (*Character, error) {
	name = strings.ToLower(name)
	if c, ok := CharacterCahceMap[name]; ok {
		if !c.IsInvalid() {
			c.Character.IsCached = true
			return &c.Character, nil
		} else {
			delete(CharacterCahceMap, name)
		}
	}
	tmp := new(Characters)
	url := fmt.Sprintf("%v%v/get/%v/character/?name.first_lower=%v&c:resolve=item,profile,faction,stat,weapon_stat,stat_history,online_status,friends,world,outfit",
		BaseURL, c.serviceID, c.namespace, name)
	if err := decode(c, url, tmp); err != nil {
		return nil, err
	}
	if len(tmp.Characters) < 1 {
		return nil, ErrCharDoesNotExist
	}
	char := tmp.Characters[0]
	char.Parent = c
	cc := &CharacterCache{}
	cc.Character = char
	cc.Cache = NewCache()
	CharacterCahceMap[name] = cc
	return &char, nil
}
