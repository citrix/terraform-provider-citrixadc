package lb

type Lbmonbindingsservicebinding struct {
	Ipaddress   string `json:"ipaddress,omitempty"`
	Monitorname string `json:"monitorname,omitempty"`
	Monsvcstate string `json:"monsvcstate,omitempty"`
	Port        int    `json:"port,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Servicetype string `json:"servicetype,omitempty"`
	Svrstate    string `json:"svrstate,omitempty"`
}
