package ssl

type Sslpkcs8 struct {
	Keyfile   string `json:"keyfile,omitempty"`
	Keyform   string `json:"keyform,omitempty"`
	Password  string `json:"password,omitempty"`
	Pkcs8file string `json:"pkcs8file,omitempty"`
}
