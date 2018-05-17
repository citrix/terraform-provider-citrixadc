package dns

type Dnsnsrec struct {
	Authtype   string `json:"authtype,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Nameserver string `json:"nameserver,omitempty"`
	Ttl        int    `json:"ttl,omitempty"`
	Type       string `json:"type,omitempty"`
}
