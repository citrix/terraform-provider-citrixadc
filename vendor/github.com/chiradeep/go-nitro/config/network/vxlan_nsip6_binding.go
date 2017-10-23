package network

type Vxlannsip6binding struct {
	Id        int    `json:"id,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}
