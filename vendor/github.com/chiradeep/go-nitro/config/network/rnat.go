package network

type Rnat struct {
	Aclname          string `json:"aclname,omitempty"`
	Connfailover     string `json:"connfailover,omitempty"`
	Natip            string `json:"natip,omitempty"`
	Natip2           string `json:"natip2,omitempty"`
	Netmask          string `json:"netmask,omitempty"`
	Network          string `json:"network,omitempty"`
	Ownergroup       string `json:"ownergroup,omitempty"`
	Redirectport     bool   `json:"redirectport,omitempty"`
	Srcippersistency string `json:"srcippersistency,omitempty"`
	Td               int    `json:"td,omitempty"`
	Useproxyport     string `json:"useproxyport,omitempty"`
}
