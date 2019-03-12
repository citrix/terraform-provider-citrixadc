package rise

type Riserhi struct {
	Hostrtgw        string `json:"hostrtgw,omitempty"`
	Ipaddress       string `json:"ipaddress,omitempty"`
	Nexthopvlan     int    `json:"nexthopvlan,omitempty"`
	Prefixlen       int    `json:"prefixlen,omitempty"`
	Vserverrhilevel string `json:"vserverrhilevel,omitempty"`
	Weight          int    `json:"weight,omitempty"`
}
