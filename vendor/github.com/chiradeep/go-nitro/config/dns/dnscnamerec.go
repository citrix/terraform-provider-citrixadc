package dns

type Dnscnamerec struct {
	Aliasname     string `json:"aliasname,omitempty"`
	Authtype      string `json:"authtype,omitempty"`
	Canonicalname string `json:"canonicalname,omitempty"`
	Ttl           int    `json:"ttl,omitempty"`
	Type          string `json:"type,omitempty"`
	Vservername   string `json:"vservername,omitempty"`
}
