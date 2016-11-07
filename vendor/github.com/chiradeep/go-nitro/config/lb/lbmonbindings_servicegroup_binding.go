package lb

type Lbmonbindingsservicegroupbinding struct {
	Boundservicegroupsvrstate string `json:"boundservicegroupsvrstate,omitempty"`
	Monitorname               string `json:"monitorname,omitempty"`
	Monstate                  string `json:"monstate,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	Servicetype               string `json:"servicetype,omitempty"`
}
