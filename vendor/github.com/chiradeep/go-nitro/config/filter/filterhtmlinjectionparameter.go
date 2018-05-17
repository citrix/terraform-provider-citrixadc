package filter

type Filterhtmlinjectionparameter struct {
	Frequency     int    `json:"frequency,omitempty"`
	Htmlsearchlen int    `json:"htmlsearchlen,omitempty"`
	Rate          int    `json:"rate,omitempty"`
	Strict        string `json:"strict,omitempty"`
}
