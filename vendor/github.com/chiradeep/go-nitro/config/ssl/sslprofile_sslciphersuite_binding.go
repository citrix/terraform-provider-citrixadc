package ssl

type Sslprofilesslciphersuitebinding struct {
	Ciphername     string `json:"ciphername,omitempty"`
	Cipherpriority int    `json:"cipherpriority,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
}
