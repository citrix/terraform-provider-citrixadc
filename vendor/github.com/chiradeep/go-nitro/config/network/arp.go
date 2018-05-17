package network

type Arp struct {
	All       bool   `json:"all,omitempty"`
	Channel   int    `json:"channel,omitempty"`
	Flags     int    `json:"flags,omitempty"`
	Ifnum     string `json:"ifnum,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Mac       string `json:"mac,omitempty"`
	Ownernode int    `json:"ownernode,omitempty"`
	State     int    `json:"state,omitempty"`
	Td        int    `json:"td,omitempty"`
	Timeout   int    `json:"timeout,omitempty"`
	Type      string `json:"type,omitempty"`
	Vlan      int    `json:"vlan,omitempty"`
}
