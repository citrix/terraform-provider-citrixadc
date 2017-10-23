package ssl

type Sslcipher struct {
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
	Ciphgrpalias    string `json:"ciphgrpalias,omitempty"`
}
