package table

import "sort"

type Player struct {
	Name    string `json:"name,omitempty"`
	Flag    string `json:"flag,omitempty"`
	Scores  []int  `json:"scores,omitempty"`
	Penalty int    `json:"penalty,omitempty"`
}

func (player *Player) Total() int {
	total := 0

	total += player.Penalty

	for _, score := range player.Scores {
		total += score
	}

	return total
}

func SortByScore(players []*Player) []*Player {
	sort.Sort(byScore(players))

	return players
}

type byScore []*Player

func (a byScore) Len() int {
	return len(a)
}

func (a byScore) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byScore) Less(i, j int) bool {
	return a[i].Total() < a[j].Total()
}
