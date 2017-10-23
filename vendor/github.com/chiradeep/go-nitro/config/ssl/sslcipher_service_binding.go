package ssl

type Sslcipherservicebinding struct {
	Ciphergroupname  string `json:"ciphergroupname,omitempty"`
	Cipheroperation  string `json:"cipheroperation,omitempty"`
	Ciphgrpals       string `json:"ciphgrpals,omitempty"`
	Service          bool   `json:"service,omitempty"`
	Servicegroup     bool   `json:"servicegroup,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
}
