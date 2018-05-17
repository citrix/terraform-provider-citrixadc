package network

type Vridnsip6binding struct {
	Flags     int    `json:"flags,omitempty"`
	Id        int    `json:"id,omitempty"`
	Ifnum     string `json:"ifnum,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
}
