package ssl

type Sslciphersslprofilebinding struct {
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
	Description     string `json:"description,omitempty"`
	Sslprofile      string `json:"sslprofile,omitempty"`
}
