package basic

type Servicegroupbindings struct {
	Ipaddress        string `json:"ipaddress,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	State            string `json:"state,omitempty"`
	Svrstate         string `json:"svrstate,omitempty"`
	Vservername      string `json:"vservername,omitempty"`
}
