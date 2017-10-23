package ssl

type Sslhsmkey struct {
	Hsmkeyname string `json:"hsmkeyname,omitempty"`
	Key        string `json:"key,omitempty"`
}
