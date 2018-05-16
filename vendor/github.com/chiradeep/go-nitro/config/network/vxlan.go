package network

type Vxlan struct {
	Dynamicrouting     string `json:"dynamicrouting,omitempty"`
	Id                 int    `json:"id,omitempty"`
	Ipv6dynamicrouting string `json:"ipv6dynamicrouting,omitempty"`
	Port               int    `json:"port,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
}
