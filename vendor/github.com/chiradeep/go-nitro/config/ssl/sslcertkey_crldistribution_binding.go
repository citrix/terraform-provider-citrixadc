package ssl

type Sslcertkeycrldistributionbinding struct {
	Ca      bool   `json:"ca,omitempty"`
	Certkey string `json:"certkey,omitempty"`
	Issuer  string `json:"issuer,omitempty"`
}
