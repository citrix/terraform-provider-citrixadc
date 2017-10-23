package ssl

type Sslciphersuite struct {
	Ciphername  string `json:"ciphername,omitempty"`
	Description string `json:"description,omitempty"`
}
