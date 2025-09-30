package ltrc

type Event struct {
	ID     int    `json:"id,omitempty"`
	Format Format `json:"format"`
}

type Racer struct {
	Event int     `json:"event,omitempty"`
	Name  string  `json:"name,omitempty"`
	Mmr   float64 `json:"mmr,omitempty"`
	Score float64 `json:"score,omitempty"`
}
