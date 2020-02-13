package ssl

type Sslhsmkey struct {
	Hsmkeyname string `json:"hsmkeyname,omitempty"`
	Hsmtype    string `json:"hsmtype,omitempty"`
	Key        string `json:"key,omitempty"`
	Keystore   string `json:"keystore,omitempty"`
	Password   string `json:"password,omitempty"`
	Serialnum  string `json:"serialnum,omitempty"`
	State      string `json:"state,omitempty"`
}
