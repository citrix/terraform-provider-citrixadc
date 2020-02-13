package ssl

type Sslcertificatechain struct {
	Certkeyname        string      `json:"certkeyname,omitempty"`
	Chaincomplete      int         `json:"chaincomplete,omitempty"`
	Chainissuer        string      `json:"chainissuer,omitempty"`
	Chainlinked        interface{} `json:"chainlinked,omitempty"`
	Chainpossiblelinks interface{} `json:"chainpossiblelinks,omitempty"`
}
