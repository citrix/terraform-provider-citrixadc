package network

type Nd6 struct {
	Channel  int    `json:"channel,omitempty"`
	Flags    int    `json:"flags,omitempty"`
	Ifnum    string `json:"ifnum,omitempty"`
	Mac      string `json:"mac,omitempty"`
	Neighbor string `json:"neighbor,omitempty"`
	State    string `json:"state,omitempty"`
	Td       int    `json:"td,omitempty"`
	Timeout  int    `json:"timeout,omitempty"`
	Vlan     int    `json:"vlan,omitempty"`
}
