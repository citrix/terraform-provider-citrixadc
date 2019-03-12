package ssl

type Sslhsmkey struct {
	Hsmkeyname string `json:"hsmkeyname,omitempty"`
	Hsmtype    string `json:"hsmtype,omitempty"`
	Key        string `json:"key,omitempty"`
	Password   string `json:"password,omitempty"`
	Serialnum  string `json:"serialnum,omitempty"`
}
