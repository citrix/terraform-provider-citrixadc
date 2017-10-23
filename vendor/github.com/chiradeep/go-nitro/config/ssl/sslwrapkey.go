package ssl

type Sslwrapkey struct {
	Password    string `json:"password,omitempty"`
	Salt        string `json:"salt,omitempty"`
	Wrapkeyname string `json:"wrapkeyname,omitempty"`
}
