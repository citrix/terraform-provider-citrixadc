package gslb

type Gslbconfig struct {
	Command    string `json:"command,omitempty"`
	Debug      bool   `json:"debug,omitempty"`
	Forcesync  string `json:"forcesync,omitempty"`
	Nowarn     bool   `json:"nowarn,omitempty"`
	Preview    bool   `json:"preview,omitempty"`
	Saveconfig bool   `json:"saveconfig,omitempty"`
}
