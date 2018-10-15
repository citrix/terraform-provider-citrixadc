package gslb

type Gslbsitegslbservicegroupmemberbinding struct {
	Ipaddress        string `json:"ipaddress,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	Sitename         string `json:"sitename,omitempty"`
	State            string `json:"state,omitempty"`
}
