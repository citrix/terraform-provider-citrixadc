package network

type Onlinkipv6prefix struct {
	Autonomusprefix          string `json:"autonomusprefix,omitempty"`
	Decrementprefixlifetimes string `json:"decrementprefixlifetimes,omitempty"`
	Depricateprefix          string `json:"depricateprefix,omitempty"`
	Ipv6prefix               string `json:"ipv6prefix,omitempty"`
	Onlinkprefix             string `json:"onlinkprefix,omitempty"`
	Prefixcurrpreferredlft   int    `json:"prefixcurrpreferredlft,omitempty"`
	Prefixcurrvalidelft      int    `json:"prefixcurrvalidelft,omitempty"`
	Prefixpreferredlifetime  int    `json:"prefixpreferredlifetime,omitempty"`
	Prefixvalidelifetime     int    `json:"prefixvalidelifetime,omitempty"`
}
