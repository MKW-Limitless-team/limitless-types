package table

import "sort"

type Group struct {
	Name    string    `json:"name,omitempty"`
	Desc    string    `json:"desc,omitempty"`
	Players []*Player `json:"players,omitempty"`
	Penalty int       `json:"penalty,omitempty"`
	Color   string    `json:"color,omitempty"`
}

func (group *Group) AddPlayer(player *Player) {
	group.Players = append(group.Players, player)
}

func (group *Group) Total() int {
	total := 0

	for _, player := range group.Players {
		total += player.Penalty

		for _, score := range player.Scores {
			total += score
		}
	}

	return total
}

func SortByTotalScore(groups []*Group) []*Group {
	sort.Sort(byTotalScore(groups))

	return groups
}

type byTotalScore []*Group

func (a byTotalScore) Len() int {
	return len(a)
}

func (a byTotalScore) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byTotalScore) Less(i, j int) bool {
	return a[i].Total() < a[j].Total()
}
