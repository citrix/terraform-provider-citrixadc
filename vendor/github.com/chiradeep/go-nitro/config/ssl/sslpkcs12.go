package ssl

type Sslpkcs12 struct {
	Certfile      string `json:"certfile,omitempty"`
	Des           bool   `json:"des,omitempty"`
	Des3          bool   `json:"des3,omitempty"`
	Export        bool   `json:"export,omitempty"`
	Import        bool   `json:"Import,omitempty"`
	Keyfile       string `json:"keyfile,omitempty"`
	Outfile       string `json:"outfile,omitempty"`
	Password      string `json:"password,omitempty"`
	Pempassphrase string `json:"pempassphrase,omitempty"`
	Pkcs12file    string `json:"pkcs12file,omitempty"`
}
