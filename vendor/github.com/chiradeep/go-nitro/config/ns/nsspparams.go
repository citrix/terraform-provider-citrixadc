package ns

type Nsspparams struct {
	Basethreshold int         `json:"basethreshold,omitempty"`
	Table0        interface{} `json:"table0,omitempty"`
	Throttle      string      `json:"throttle,omitempty"`
}
