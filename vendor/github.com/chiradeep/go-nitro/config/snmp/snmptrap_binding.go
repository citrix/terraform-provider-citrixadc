package snmp

type Snmptrapbinding struct {
	Td              int    `json:"td,omitempty"`
	Trapclass       string `json:"trapclass,omitempty"`
	Trapdestination string `json:"trapdestination,omitempty"`
	Version         string `json:"version,omitempty"`
}
