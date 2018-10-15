package dns

type Dnsaddrec struct {
	Authtype    string `json:"authtype,omitempty"`
	Ecssubnet   string `json:"ecssubnet,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Nodeid      int    `json:"nodeid,omitempty"`
	Ttl         int    `json:"ttl,omitempty"`
	Type        string `json:"type,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}
