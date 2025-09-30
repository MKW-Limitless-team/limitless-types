package ltrc

import (
	"sort"
)

type Placement struct {
	ID           int    `json:"id,omitempty"`
	Track        string `json:"track,omitempty"`
	DiscordID    string `json:"discord_id,omitempty"`
	Minutes      int    `json:"minutes,omitempty"`
	Seconds      int    `json:"seconds,omitempty"`
	Milliseconds int    `json:"milliseconds,omitempty"`
	Character    string `json:"character,omitempty"`
	Vehicle      string `json:"vehicle,omitempty"`
	DriftType    string `json:"drift_type,omitempty"`
	Category     string `json:"category,omitempty"`
	Url          string `json:"url,omitempty"`
	CRC          uint32 `json:"crc,omitempty"`
	Approved     bool   `json:"approved,omitempty"`
}

type byTime []*Placement

func (b byTime) Len() int {
	return len(b)
}

func (b byTime) Less(i int, j int) bool {
	time1 := b[i]
	time2 := b[j]

	if time1.Minutes > time2.Minutes {
		return true
	} else if time1.Seconds > time2.Seconds {
		return true
	} else if time1.Milliseconds > time2.Milliseconds {
		return true
	}

	return false
}

func (b byTime) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func SortByTime(placements []*Placement) []*Placement {
	sort.Sort(byTime(placements))

	return placements
}
