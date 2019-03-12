package network

type Vridinterfacebinding struct {
	Flags int    `json:"flags,omitempty"`
	Id    int    `json:"id,omitempty"`
	Ifnum string `json:"ifnum,omitempty"`
	Vlan  int    `json:"vlan,omitempty"`
}
