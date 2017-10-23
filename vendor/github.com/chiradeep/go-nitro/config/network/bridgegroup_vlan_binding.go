package network

type Bridgegroupvlanbinding struct {
	Id   int  `json:"id,omitempty"`
	Rnat bool `json:"rnat,omitempty"`
	Vlan int  `json:"vlan,omitempty"`
}
