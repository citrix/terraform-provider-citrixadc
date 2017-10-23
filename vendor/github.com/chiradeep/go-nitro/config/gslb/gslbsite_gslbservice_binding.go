package gslb

type Gslbsitegslbservicebinding struct {
	Cnameentry  string `json:"cnameentry,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Port        int    `json:"port,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Servicetype string `json:"servicetype,omitempty"`
	Sitename    string `json:"sitename,omitempty"`
	State       string `json:"state,omitempty"`
}
