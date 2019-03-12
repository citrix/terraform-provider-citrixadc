package gslb

type Gslbdomaingslbservicegroupmemberbinding struct {
	Gslbthreshold    int    `json:"gslbthreshold,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Name             string `json:"name,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	Svreffgslbstate  string `json:"svreffgslbstate,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}
