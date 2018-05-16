package snmp

type Snmptrapsnmpuserbinding struct {
	Securitylevel   string `json:"securitylevel,omitempty"`
	Td              int    `json:"td,omitempty"`
	Trapclass       string `json:"trapclass,omitempty"`
	Trapdestination string `json:"trapdestination,omitempty"`
	Username        string `json:"username,omitempty"`
	Version         string `json:"version,omitempty"`
}
