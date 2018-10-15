package network

type Vxlan struct {
	Dynamicrouting     string `json:"dynamicrouting,omitempty"`
	Id                 int    `json:"id,omitempty"`
	Innervlantagging   string `json:"innervlantagging,omitempty"`
	Ipv6dynamicrouting string `json:"ipv6dynamicrouting,omitempty"`
	Partitionname      string `json:"partitionname,omitempty"`
	Port               int    `json:"port,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Type               string `json:"type,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
}
