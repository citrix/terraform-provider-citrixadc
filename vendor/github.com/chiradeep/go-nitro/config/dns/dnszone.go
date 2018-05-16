package dns

type Dnszone struct {
	Dnssecoffload string      `json:"dnssecoffload,omitempty"`
	Flags         int         `json:"flags,omitempty"`
	Keyname       interface{} `json:"keyname,omitempty"`
	Nsec          string      `json:"nsec,omitempty"`
	Proxymode     string      `json:"proxymode,omitempty"`
	Type          string      `json:"type,omitempty"`
	Zonename      string      `json:"zonename,omitempty"`
}
