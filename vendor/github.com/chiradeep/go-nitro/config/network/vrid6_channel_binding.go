package network

type Vrid6channelbinding struct {
	Flags int    `json:"flags,omitempty"`
	Id    int    `json:"id,omitempty"`
	Ifnum string `json:"ifnum,omitempty"`
	Vlan  int    `json:"vlan,omitempty"`
}
