package dns

type Dnsnsecrec struct {
	Ecssubnet string      `json:"ecssubnet,omitempty"`
	Hostname  string      `json:"hostname,omitempty"`
	Nextnsec  string      `json:"nextnsec,omitempty"`
	Nextrecs  interface{} `json:"nextrecs,omitempty"`
	Ttl       int         `json:"ttl,omitempty"`
	Type      string      `json:"type,omitempty"`
}
