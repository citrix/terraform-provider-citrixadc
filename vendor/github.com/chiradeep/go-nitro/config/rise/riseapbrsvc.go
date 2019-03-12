package rise

type Riseapbrsvc struct {
	Name        string `json:"name,omitempty"`
	Nexthopip   string `json:"nexthopip,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Risesvctype string `json:"risesvctype,omitempty"`
	Serverip    string `json:"serverip,omitempty"`
	Serverport  int    `json:"serverport,omitempty"`
	Vlan        int    `json:"vlan,omitempty"`
}
