package vpn

type Vpnnexthopserver struct {
	Name        string `json:"name,omitempty"`
	Nexthopip   string `json:"nexthopip,omitempty"`
	Nexthopport int    `json:"nexthopport,omitempty"`
	Secure      string `json:"secure,omitempty"`
}
