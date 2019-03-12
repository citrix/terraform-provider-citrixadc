package network

type Vridnsipbinding struct {
	Flags     int    `json:"flags,omitempty"`
	Id        int    `json:"id,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
}
