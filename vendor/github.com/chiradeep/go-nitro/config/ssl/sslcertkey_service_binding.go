package ssl

type Sslcertkeyservicebinding struct {
	Ca               bool   `json:"ca,omitempty"`
	Certkey          string `json:"certkey,omitempty"`
	Crlcheck         string `json:"crlcheck,omitempty"`
	Data             int    `json:"data,omitempty"`
	Service          bool   `json:"service,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Version          int    `json:"version,omitempty"`
}
