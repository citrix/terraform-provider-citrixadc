package vpn

type Vpnnexthopserver struct {
	Name           string `json:"name,omitempty"`
	Nexthopfqdn    string `json:"nexthopfqdn,omitempty"`
	Nexthopip      string `json:"nexthopip,omitempty"`
	Nexthopport    int    `json:"nexthopport,omitempty"`
	Resaddresstype string `json:"resaddresstype,omitempty"`
	Secure         string `json:"secure,omitempty"`
}
