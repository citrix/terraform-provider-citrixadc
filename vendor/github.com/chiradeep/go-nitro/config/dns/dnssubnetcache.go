package dns

type Dnssubnetcache struct {
	All       bool        `json:"all,omitempty"`
	Ecssubnet string      `json:"ecssubnet,omitempty"`
	Hostname  string      `json:"hostname,omitempty"`
	Nextrecs  interface{} `json:"nextrecs,omitempty"`
	Nodeid    int         `json:"nodeid,omitempty"`
}
