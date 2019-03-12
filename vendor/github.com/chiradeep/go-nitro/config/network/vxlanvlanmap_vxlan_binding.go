package network

type Vxlanvlanmapvxlanbinding struct {
	Name  string      `json:"name,omitempty"`
	Vlan  interface{} `json:"vlan,omitempty"`
	Vxlan int         `json:"vxlan,omitempty"`
}
