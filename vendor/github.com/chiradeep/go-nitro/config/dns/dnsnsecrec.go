package dns

type Dnsnsecrec struct {
	Hostname string      `json:"hostname,omitempty"`
	Nextnsec string      `json:"nextnsec,omitempty"`
	Nextrecs interface{} `json:"nextrecs,omitempty"`
	Type     string      `json:"type,omitempty"`
}
