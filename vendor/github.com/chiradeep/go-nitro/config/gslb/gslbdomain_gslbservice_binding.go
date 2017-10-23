package gslb

type Gslbdomaingslbservicebinding struct {
	Cnameentry       string `json:"cnameentry,omitempty"`
	Cumulativeweight int    `json:"cumulativeweight,omitempty"`
	Dynamicconfwt    int    `json:"dynamicconfwt,omitempty"`
	Gslbthreshold    int    `json:"gslbthreshold,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Name             string `json:"name,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	State            string `json:"state,omitempty"`
	Svreffgslbstate  string `json:"svreffgslbstate,omitempty"`
	Vservername      string `json:"vservername,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}
