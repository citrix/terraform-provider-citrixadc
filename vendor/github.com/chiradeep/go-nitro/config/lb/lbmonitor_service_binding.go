package lb

type Lbmonitorservicebinding struct {
	Dupstate         string `json:"dup_state,omitempty"`
	Dupweight        int    `json:"dup_weight,omitempty"`
	Monitorname      string `json:"monitorname,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}
