package ipsecalg

type Ipsecalgsession struct {
	Destip      string `json:"destip,omitempty"`
	Destipalg   string `json:"destip_alg,omitempty"`
	Natip       string `json:"natip,omitempty"`
	Natipalg    string `json:"natip_alg,omitempty"`
	Sourceip    string `json:"sourceip,omitempty"`
	Sourceipalg string `json:"sourceip_alg,omitempty"`
	Spiin       int    `json:"spiin,omitempty"`
	Spiout      int    `json:"spiout,omitempty"`
}
