package network

type Bridgetable struct {
	Bridgeage int    `json:"bridgeage,omitempty"`
	Channel   int    `json:"channel,omitempty"`
	Ifnum     string `json:"ifnum,omitempty"`
	Mac       string `json:"mac,omitempty"`
	Vlan      int    `json:"vlan,omitempty"`
}
