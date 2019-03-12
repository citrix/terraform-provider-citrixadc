package cluster

type Clusterpropstatus struct {
	Cmdstrs          string `json:"cmdstrs,omitempty"`
	Nodeid           int    `json:"nodeid,omitempty"`
	Numpropcmdfailed int    `json:"numpropcmdfailed,omitempty"`
}
