package ssl

type Sslciphersslvserverbinding struct {
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
	Vserver         bool   `json:"vserver,omitempty"`
	Vservername     string `json:"vservername,omitempty"`
}
