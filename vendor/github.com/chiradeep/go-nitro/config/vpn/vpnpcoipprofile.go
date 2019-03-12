package vpn

type Vpnpcoipprofile struct {
	Conserverurl       string `json:"conserverurl,omitempty"`
	Icvverification    string `json:"icvverification,omitempty"`
	Name               string `json:"name,omitempty"`
	Sessionidletimeout int    `json:"sessionidletimeout,omitempty"`
}
