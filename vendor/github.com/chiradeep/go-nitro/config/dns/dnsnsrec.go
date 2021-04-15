package dns

type Dnsnsrec struct {
	Authtype   string `json:"authtype,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Ecssubnet  string `json:"ecssubnet,omitempty"`
	Nameserver string `json:"nameserver,omitempty"`
	Nodeid     int    `json:"nodeid,omitempty"`
	Ttl        int    `json:"ttl,omitempty"`
	Type       string `json:"type,omitempty"`
}
