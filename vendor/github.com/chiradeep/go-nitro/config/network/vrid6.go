package network

type Vrid6 struct {
	All       bool   `json:"all,omitempty"`
	Flags     int    `json:"flags,omitempty"`
	Id        int    `json:"id,omitempty"`
	Ifaces    string `json:"ifaces,omitempty"`
	Ifnum     string `json:"ifnum,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Priority  int    `json:"priority,omitempty"`
	State     int    `json:"state,omitempty"`
	Type      string `json:"type,omitempty"`
}
