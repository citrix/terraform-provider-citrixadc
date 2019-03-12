package network

type Rnatsession struct {
	Aclname string `json:"aclname,omitempty"`
	Natip   string `json:"natip,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Network string `json:"network,omitempty"`
}
