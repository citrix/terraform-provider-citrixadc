package lb

type Lbvserverservicegroupmemberbinding struct {
	Cookieipport      string `json:"cookieipport,omitempty"`
	Cookiename        string `json:"cookiename,omitempty"`
	Curstate          string `json:"curstate,omitempty"`
	Dynamicweight     int    `json:"dynamicweight,omitempty"`
	Ipv46             string `json:"ipv46,omitempty"`
	Name              string `json:"name,omitempty"`
	Port              int    `json:"port,omitempty"`
	Preferredlocation string `json:"preferredlocation,omitempty"`
	Servicegroupname  string `json:"servicegroupname,omitempty"`
	Servicetype       string `json:"servicetype,omitempty"`
	Vserverid         string `json:"vserverid,omitempty"`
	Weight            int    `json:"weight,omitempty"`
}
