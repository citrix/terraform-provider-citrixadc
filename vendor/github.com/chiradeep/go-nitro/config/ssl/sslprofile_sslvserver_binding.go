package ssl

type Sslprofilesslvserverbinding struct {
	Cipherpriority int    `json:"cipherpriority,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	Servicename    string `json:"servicename,omitempty"`
}
