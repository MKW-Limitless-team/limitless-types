package ltrc

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/MKW-Limitless-team/limitless-types/wwfc"
)

type PlayerData struct {
	Name      string `json:"name,omitempty"`
	ProfileID uint64 `json:"profileID,omitempty"`
	DiscordID string `json:"discordID,omitempty"`
	Mmr       *int64 `json:"mmr,omitempty"`
}

func (playerData *PlayerData) GetFC() string {
	fc := strconv.Itoa(int(wwfc.PidToFC(playerData.ProfileID)))
	for len(fc) != 12 {
		fc = "0" + fc
	}
	return fmt.Sprintf("%s-%s-%s", fc[:4], fc[4:8], fc[8:12])
}

type byMmr []*PlayerData

func (b byMmr) Len() int {
	return len(b)
}

func (b byMmr) Less(i int, j int) bool {
	return *b[i].Mmr > *b[j].Mmr
}

func (b byMmr) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func SortByMMR(players []*PlayerData) []*PlayerData {
	sort.Sort(byMmr(players))

	return players
}
