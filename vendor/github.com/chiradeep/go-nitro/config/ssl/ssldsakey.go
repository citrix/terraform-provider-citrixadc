package ssl

type Ssldsakey struct {
	Bits     int    `json:"bits,omitempty"`
	Des      bool   `json:"des,omitempty"`
	Des3     bool   `json:"des3,omitempty"`
	Keyfile  string `json:"keyfile,omitempty"`
	Keyform  string `json:"keyform,omitempty"`
	Password string `json:"password,omitempty"`
}
