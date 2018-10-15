package ssl

type Sslecdsakey struct {
	Aes256   bool   `json:"aes256,omitempty"`
	Curve    string `json:"curve,omitempty"`
	Des      bool   `json:"des,omitempty"`
	Des3     bool   `json:"des3,omitempty"`
	Keyfile  string `json:"keyfile,omitempty"`
	Keyform  string `json:"keyform,omitempty"`
	Password string `json:"password,omitempty"`
	Pkcs8    bool   `json:"pkcs8,omitempty"`
}
