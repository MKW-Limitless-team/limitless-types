package ltrc

type Participant struct {
	Name string `json:"name,omitempty"`
	Mmr  int64  `json:"mmr,omitempty"`
}
