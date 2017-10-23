package vpn

type Vpnvserverintranetipbinding struct {
	Acttype    int    `json:"acttype,omitempty"`
	Intranetip string `json:"intranetip,omitempty"`
	Map        string `json:"map,omitempty"`
	Name       string `json:"name,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
}
