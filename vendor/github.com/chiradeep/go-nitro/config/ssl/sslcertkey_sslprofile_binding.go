package ssl

type Sslcertkeysslprofilebinding struct {
	Ca         bool   `json:"ca,omitempty"`
	Certkey    string `json:"certkey,omitempty"`
	Sslprofile string `json:"sslprofile,omitempty"`
}
