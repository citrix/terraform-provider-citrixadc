package dns

type Dnssrvrec struct {
	Authtype  string `json:"authtype,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Ecssubnet string `json:"ecssubnet,omitempty"`
	Nodeid    int    `json:"nodeid,omitempty"`
	Port      int    `json:"port,omitempty"`
	Priority  int    `json:"priority,omitempty"`
	Target    string `json:"target,omitempty"`
	Ttl       int    `json:"ttl,omitempty"`
	Type      string `json:"type,omitempty"`
	Weight    int    `json:"weight,omitempty"`
}
