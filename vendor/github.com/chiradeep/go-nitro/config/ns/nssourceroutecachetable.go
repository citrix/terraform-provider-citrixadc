package ns

type Nssourceroutecachetable struct {
	Interface string `json:"Interface,omitempty"`
	Sourceip  string `json:"sourceip,omitempty"`
	Sourcemac string `json:"sourcemac,omitempty"`
	Vlan      int    `json:"vlan,omitempty"`
}
