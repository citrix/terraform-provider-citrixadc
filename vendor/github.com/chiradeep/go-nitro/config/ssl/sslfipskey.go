package ssl

type Sslfipskey struct {
	Curve       string `json:"curve,omitempty"`
	Exponent    string `json:"exponent,omitempty"`
	Fipskeyname string `json:"fipskeyname,omitempty"`
	Inform      string `json:"inform,omitempty"`
	Iv          string `json:"iv,omitempty"`
	Key         string `json:"key,omitempty"`
	Keytype     string `json:"keytype,omitempty"`
	Modulus     int    `json:"modulus,omitempty"`
	Size        int    `json:"size,omitempty"`
	Wrapkeyname string `json:"wrapkeyname,omitempty"`
}
