package system

type Systemextramgmtcpu struct {
	Configuredstate string `json:"configuredstate,omitempty"`
	Effectivestate  string `json:"effectivestate,omitempty"`
	Nodeid          int    `json:"nodeid,omitempty"`
}
