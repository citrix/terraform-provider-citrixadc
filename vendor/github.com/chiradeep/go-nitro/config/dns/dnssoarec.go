package dns

type Dnssoarec struct {
	Authtype     string `json:"authtype,omitempty"`
	Contact      string `json:"contact,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Ecssubnet    string `json:"ecssubnet,omitempty"`
	Expire       int    `json:"expire,omitempty"`
	Minimum      int    `json:"minimum,omitempty"`
	Nodeid       int    `json:"nodeid,omitempty"`
	Originserver string `json:"originserver,omitempty"`
	Refresh      int    `json:"refresh,omitempty"`
	Retry        int    `json:"retry,omitempty"`
	Serial       int    `json:"serial,omitempty"`
	Ttl          int    `json:"ttl,omitempty"`
	Type         string `json:"type,omitempty"`
}
