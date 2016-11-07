package lb

type Lbwlm struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Katimeout int    `json:"katimeout,omitempty"`
	Lbuid     string `json:"lbuid,omitempty"`
	Port      int    `json:"port,omitempty"`
	Secure    string `json:"secure,omitempty"`
	State     string `json:"state,omitempty"`
	Wlmname   string `json:"wlmname,omitempty"`
}
