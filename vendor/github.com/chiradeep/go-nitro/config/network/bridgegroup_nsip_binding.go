package network

type Bridgegroupnsipbinding struct {
	Id        int    `json:"id,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Rnat      bool   `json:"rnat,omitempty"`
	Td        int    `json:"td,omitempty"`
}
