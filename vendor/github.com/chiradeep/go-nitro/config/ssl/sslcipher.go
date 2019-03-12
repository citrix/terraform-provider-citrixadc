package ssl

type Sslcipher struct {
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
	Ciphgrpalias    string `json:"ciphgrpalias,omitempty"`
	Sslprofile      string `json:"sslprofile,omitempty"`
}
