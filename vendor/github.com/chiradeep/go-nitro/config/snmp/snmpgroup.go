package snmp

type Snmpgroup struct {
	Name          string `json:"name,omitempty"`
	Readviewname  string `json:"readviewname,omitempty"`
	Securitylevel string `json:"securitylevel,omitempty"`
	Status        string `json:"status,omitempty"`
	Storagetype   string `json:"storagetype,omitempty"`
}
