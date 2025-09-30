package ltrc

type Season struct {
	Name         string      `json:"name,omitempty"`
	Active       bool        `json:"active,omitempty"`
	Participants Participant `json:"participants"`
}
