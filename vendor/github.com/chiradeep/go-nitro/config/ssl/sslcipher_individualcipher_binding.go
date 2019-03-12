package ssl

type Sslcipherindividualcipherbinding struct {
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
	Description     string `json:"description,omitempty"`
}
