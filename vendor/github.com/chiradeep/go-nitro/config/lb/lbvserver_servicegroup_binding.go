package lb

type Lbvserverservicegroupbinding struct {
	Name             string `json:"name,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}
