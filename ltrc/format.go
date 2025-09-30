package ltrc

type Format struct {
	Name          string     `json:"name,omitempty"`
	ScalingFactor int        `json:"scaling_factor,omitempty"`
	TeamSize      int        `json:"team_size,omitempty"`
	Modifiers     []Modifier `json:"modifiers,omitempty"`
}

type Modifier struct {
	Name     string  `json:"name,omitempty"`
	Increase float64 `json:"increase,omitempty"`
	Decrease float64 `json:"decrease,omitempty"`
}
