package dns

type Dnsptrrec struct {
	Authtype      string `json:"authtype,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Reversedomain string `json:"reversedomain,omitempty"`
	Ttl           int    `json:"ttl,omitempty"`
	Type          string `json:"type,omitempty"`
}
