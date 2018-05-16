package ns

type Nsdhcpparams struct {
	Dhcpclient string `json:"dhcpclient,omitempty"`
	Hostrtgw   string `json:"hostrtgw,omitempty"`
	Ipaddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Running    bool   `json:"running,omitempty"`
	Saveroute  string `json:"saveroute,omitempty"`
}
