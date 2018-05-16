package ns

type Nssimpleacl struct {
	Aclaction   string `json:"aclaction,omitempty"`
	Aclname     string `json:"aclname,omitempty"`
	Destport    int    `json:"destport,omitempty"`
	Estsessions bool   `json:"estsessions,omitempty"`
	Hits        int    `json:"hits,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Srcip       string `json:"srcip,omitempty"`
	Td          int    `json:"td,omitempty"`
	Ttl         int    `json:"ttl,omitempty"`
}
