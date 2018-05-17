package ns

type Nssimpleacl6 struct {
	Aclaction   string `json:"aclaction,omitempty"`
	Aclname     string `json:"aclname,omitempty"`
	Destport    int    `json:"destport,omitempty"`
	Estsessions bool   `json:"estsessions,omitempty"`
	Hits        int    `json:"hits,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Srcipv6     string `json:"srcipv6,omitempty"`
	Td          int    `json:"td,omitempty"`
	Ttl         int    `json:"ttl,omitempty"`
}
