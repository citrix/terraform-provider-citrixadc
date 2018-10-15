package lb

type Lbvserverservicebinding struct {
	Cookieipport      string `json:"cookieipport,omitempty"`
	Curstate          string `json:"curstate,omitempty"`
	Dynamicweight     int    `json:"dynamicweight,omitempty"`
	Ipv46             string `json:"ipv46,omitempty"`
	Name              string `json:"name,omitempty"`
	Port              int    `json:"port,omitempty"`
	Preferredlocation string `json:"preferredlocation,omitempty"`
	Servicegroupname  string `json:"servicegroupname,omitempty"`
	Servicename       string `json:"servicename,omitempty"`
	Servicetype       string `json:"servicetype,omitempty"`
	Vserverid         string `json:"vserverid,omitempty"`
	Vsvrbindsvcip     string `json:"vsvrbindsvcip,omitempty"`
	Vsvrbindsvcport   int    `json:"vsvrbindsvcport,omitempty"`
	Weight            int    `json:"weight,omitempty"`
}
