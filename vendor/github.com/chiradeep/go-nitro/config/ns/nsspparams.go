package ns

type Nsspparams struct {
	Basethreshold int         `json:"basethreshold,omitempty"`
	Builtin       interface{} `json:"builtin,omitempty"`
	Feature       string      `json:"feature,omitempty"`
	Table0        interface{} `json:"table0,omitempty"`
	Throttle      string      `json:"throttle,omitempty"`
}
