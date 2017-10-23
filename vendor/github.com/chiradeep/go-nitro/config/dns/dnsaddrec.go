package dns

type Dnsaddrec struct {
	Authtype    string `json:"authtype,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Ttl         int    `json:"ttl,omitempty"`
	Type        string `json:"type,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}
