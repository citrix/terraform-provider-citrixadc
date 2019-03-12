package network

type Bridgetable struct {
	Bridgeage    int    `json:"bridgeage,omitempty"`
	Channel      int    `json:"channel,omitempty"`
	Controlplane bool   `json:"controlplane,omitempty"`
	Devicevlan   int    `json:"devicevlan,omitempty"`
	Flags        int    `json:"flags,omitempty"`
	Ifnum        string `json:"ifnum,omitempty"`
	Mac          string `json:"mac,omitempty"`
	Nodeid       int    `json:"nodeid,omitempty"`
	Type         string `json:"type,omitempty"`
	Vlan         int    `json:"vlan,omitempty"`
	Vni          int    `json:"vni,omitempty"`
	Vtep         string `json:"vtep,omitempty"`
	Vxlan        int    `json:"vxlan,omitempty"`
}
