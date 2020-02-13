package basic

type Member struct {
	Ip     string `json:"ip,omitempty"`
	Port   int    `json:"port,omitempty"`
	Weight int    `json:"weight,omitempty"`
}

type Servicegroupservicegroupmemberlistbinding struct {
	Servicegroupname string   `json:"servicegroupname,omitempty"`
	Members          []Member `json:"members,omitempty"`
	Failedmembers    []Member `json:"failedmembers,omitempty"`
}
