package snmp

type Snmptrap struct {
	Communityname   string `json:"communityname,omitempty"`
	Destport        int    `json:"destport,omitempty"`
	Severity        string `json:"severity,omitempty"`
	Srcip           string `json:"srcip,omitempty"`
	Trapclass       string `json:"trapclass,omitempty"`
	Trapdestination string `json:"trapdestination,omitempty"`
	Version         string `json:"version,omitempty"`
}
