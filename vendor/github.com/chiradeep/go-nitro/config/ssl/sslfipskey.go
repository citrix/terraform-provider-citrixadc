package ssl

type Sslfipskey struct {
	Exponent    string `json:"exponent,omitempty"`
	Fipskeyname string `json:"fipskeyname,omitempty"`
	Inform      string `json:"inform,omitempty"`
	Iv          string `json:"iv,omitempty"`
	Key         string `json:"key,omitempty"`
	Modulus     int    `json:"modulus,omitempty"`
	Size        int    `json:"size,omitempty"`
	Wrapkeyname string `json:"wrapkeyname,omitempty"`
}
