package network

type Rnat struct {
	Aclname      string `json:"aclname,omitempty"`
	Natip        string `json:"natip,omitempty"`
	Natip2       string `json:"natip2,omitempty"`
	Netmask      string `json:"netmask,omitempty"`
	Network      string `json:"network,omitempty"`
	Redirectport bool   `json:"redirectport,omitempty"`
	Td           int    `json:"td,omitempty"`
}
