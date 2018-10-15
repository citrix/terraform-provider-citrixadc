package ssl

type Sslcipherservicegroupbinding struct {
	Ciphergroupname  string `json:"ciphergroupname,omitempty"`
	Cipheroperation  string `json:"cipheroperation,omitempty"`
	Cipherpriority   int    `json:"cipherpriority,omitempty"`
	Ciphgrpals       string `json:"ciphgrpals,omitempty"`
	Service          bool   `json:"service,omitempty"`
	Servicegroup     bool   `json:"servicegroup,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
}
