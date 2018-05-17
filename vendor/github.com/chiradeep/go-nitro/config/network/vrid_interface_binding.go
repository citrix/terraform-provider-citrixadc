package network

type Vridinterfacebinding struct {
	Flags  int    `json:"flags,omitempty"`
	Id     int    `json:"id,omitempty"`
	Ifaces string `json:"ifaces,omitempty"`
	Ifnum  string `json:"ifnum,omitempty"`
	Vlan   int    `json:"vlan,omitempty"`
}
