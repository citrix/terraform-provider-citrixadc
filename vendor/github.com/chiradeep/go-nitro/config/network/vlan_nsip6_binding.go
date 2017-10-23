package network

type Vlannsip6binding struct {
	Id        int    `json:"id,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Td        int    `json:"td,omitempty"`
}
