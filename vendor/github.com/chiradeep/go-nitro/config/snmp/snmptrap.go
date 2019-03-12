package snmp

type Snmptrap struct {
	Allpartitions   string `json:"allpartitions,omitempty"`
	Communityname   string `json:"communityname,omitempty"`
	Destport        int    `json:"destport,omitempty"`
	Severity        string `json:"severity,omitempty"`
	Srcip           string `json:"srcip,omitempty"`
	Td              int    `json:"td,omitempty"`
	Trapclass       string `json:"trapclass,omitempty"`
	Trapdestination string `json:"trapdestination,omitempty"`
	Version         string `json:"version,omitempty"`
}
