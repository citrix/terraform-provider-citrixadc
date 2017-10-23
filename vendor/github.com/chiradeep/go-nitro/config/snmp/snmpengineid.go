package snmp

type Snmpengineid struct {
	Defaultengineid string `json:"defaultengineid,omitempty"`
	Engineid        string `json:"engineid,omitempty"`
	Ownernode       int    `json:"ownernode,omitempty"`
}
