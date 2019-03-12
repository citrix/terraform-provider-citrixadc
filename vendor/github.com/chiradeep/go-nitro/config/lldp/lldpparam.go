package lldp

type Lldpparam struct {
	Holdtimetxmult int    `json:"holdtimetxmult,omitempty"`
	Mode           string `json:"mode,omitempty"`
	Timer          int    `json:"timer,omitempty"`
}
