package cloud

type Cloudprofile struct {
	Boundservicegroupsvctype string `json:"boundservicegroupsvctype,omitempty"`
	Delay                    int    `json:"delay,omitempty"`
	Graceful                 string `json:"graceful,omitempty"`
	Ipaddress                string `json:"ipaddress,omitempty"`
	Name                     string `json:"name,omitempty"`
	Port                     int    `json:"port,omitempty"`
	Servicegroupname         string `json:"servicegroupname,omitempty"`
	Servicetype              string `json:"servicetype,omitempty"`
	Type                     string `json:"type,omitempty"`
	Vservername              string `json:"vservername,omitempty"`
	Vsvrbindsvcport          int    `json:"vsvrbindsvcport,omitempty"`
}
