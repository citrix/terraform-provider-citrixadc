package network

type Vlannsipbinding struct {
	Id         int    `json:"id,omitempty"`
	Ipaddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Td         int    `json:"td,omitempty"`
}
