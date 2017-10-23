package dns

type Dnszonedomainbinding struct {
	Domain   string      `json:"domain,omitempty"`
	Nextrecs interface{} `json:"nextrecs,omitempty"`
	Zonename string      `json:"zonename,omitempty"`
}
