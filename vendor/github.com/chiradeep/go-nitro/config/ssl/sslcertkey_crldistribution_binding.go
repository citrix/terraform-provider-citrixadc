package ssl

type Sslcertkeycrldistributionbinding struct {
	Ca       bool   `json:"ca,omitempty"`
	Certkey  string `json:"certkey,omitempty"`
	Crlcheck string `json:"crlcheck,omitempty"`
	Issuer   string `json:"issuer,omitempty"`
}
