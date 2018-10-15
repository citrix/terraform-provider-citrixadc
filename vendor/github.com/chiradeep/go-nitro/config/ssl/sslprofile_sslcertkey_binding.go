package ssl

type Sslprofilesslcertkeybinding struct {
	Cipherpriority int    `json:"cipherpriority,omitempty"`
	Name           string `json:"name,omitempty"`
	Sslicacertkey  string `json:"sslicacertkey,omitempty"`
}
