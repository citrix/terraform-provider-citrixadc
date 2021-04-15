package dns

type Dnstxtrec struct {
	Authtype  string      `json:"authtype,omitempty"`
	Domain    string      `json:"domain,omitempty"`
	Ecssubnet string      `json:"ecssubnet,omitempty"`
	Nodeid    int         `json:"nodeid,omitempty"`
	Recordid  int         `json:"recordid,omitempty"`
	String    interface{} `json:"String,omitempty"`
	Ttl       int         `json:"ttl,omitempty"`
	Type      string      `json:"type,omitempty"`
}
