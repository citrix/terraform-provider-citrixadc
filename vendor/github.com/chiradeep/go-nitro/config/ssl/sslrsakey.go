package ssl

type Sslrsakey struct {
	Aes256   bool   `json:"aes256,omitempty"`
	Bits     int    `json:"bits,omitempty"`
	Des      bool   `json:"des,omitempty"`
	Des3     bool   `json:"des3,omitempty"`
	Exponent string `json:"exponent,omitempty"`
	Keyfile  string `json:"keyfile,omitempty"`
	Keyform  string `json:"keyform,omitempty"`
	Password string `json:"password,omitempty"`
	Pkcs8    bool   `json:"pkcs8,omitempty"`
}
