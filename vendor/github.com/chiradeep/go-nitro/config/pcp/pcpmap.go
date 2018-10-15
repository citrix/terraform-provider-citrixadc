package pcp

type Pcpmap struct {
	Nattype     string `json:"nattype,omitempty"`
	Pcpaddr     int    `json:"pcpaddr,omitempty"`
	Pcpdstip    string `json:"pcpdstip,omitempty"`
	Pcpdstport  int    `json:"pcpdstport,omitempty"`
	Pcplifetime int    `json:"pcplifetime,omitempty"`
	Pcpnatip    string `json:"pcpnatip,omitempty"`
	Pcpnatport  int    `json:"pcpnatport,omitempty"`
	Pcpnounce   int    `json:"pcpnounce,omitempty"`
	Pcpprotocol string `json:"pcpprotocol,omitempty"`
	Pcprefcnt   int    `json:"pcprefcnt,omitempty"`
	Pcpsrcip    string `json:"pcpsrcip,omitempty"`
	Pcpsrcport  int    `json:"pcpsrcport,omitempty"`
	Subscrip    string `json:"subscrip,omitempty"`
}
