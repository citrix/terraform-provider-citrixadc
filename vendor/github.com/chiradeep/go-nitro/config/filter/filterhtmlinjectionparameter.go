package filter

type Filterhtmlinjectionparameter struct {
	Builtin       interface{} `json:"builtin,omitempty"`
	Feature       string      `json:"feature,omitempty"`
	Frequency     int         `json:"frequency,omitempty"`
	Htmlsearchlen int         `json:"htmlsearchlen,omitempty"`
	Rate          int         `json:"rate,omitempty"`
	Strict        string      `json:"strict,omitempty"`
}
