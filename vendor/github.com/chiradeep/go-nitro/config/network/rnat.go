package network

type Rnat struct {
	Aclname          string `json:"aclname,omitempty"`
	Connfailover     string `json:"connfailover,omitempty"`
	Name             string `json:"name,omitempty"`
	Natip            string `json:"natip,omitempty"`
	Netmask          string `json:"netmask,omitempty"`
	Network          string `json:"network,omitempty"`
	Newname          string `json:"newname,omitempty"`
	Ownergroup       string `json:"ownergroup,omitempty"`
	Redirectport     int    `json:"redirectport,omitempty"`
	Srcippersistency string `json:"srcippersistency,omitempty"`
	Td               int    `json:"td,omitempty"`
	Useproxyport     string `json:"useproxyport,omitempty"`
}
