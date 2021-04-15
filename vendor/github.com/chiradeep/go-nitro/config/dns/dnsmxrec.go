package dns

type Dnsmxrec struct {
	Authtype  string `json:"authtype,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Ecssubnet string `json:"ecssubnet,omitempty"`
	Mx        string `json:"mx,omitempty"`
	Nodeid    int    `json:"nodeid,omitempty"`
	Pref      int    `json:"pref,omitempty"`
	Ttl       int    `json:"ttl,omitempty"`
	Type      string `json:"type,omitempty"`
}
