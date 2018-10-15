package snmp

type Snmpmib struct {
	Contact     string `json:"contact,omitempty"`
	Customid    string `json:"customid,omitempty"`
	Location    string `json:"location,omitempty"`
	Name        string `json:"name,omitempty"`
	Ownernode   int    `json:"ownernode,omitempty"`
	Sysdesc     string `json:"sysdesc,omitempty"`
	Sysoid      string `json:"sysoid,omitempty"`
	Sysservices int    `json:"sysservices,omitempty"`
	Sysuptime   int    `json:"sysuptime,omitempty"`
}
