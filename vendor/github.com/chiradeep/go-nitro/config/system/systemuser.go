package system

type Systemuser struct {
	Encrypted           bool   `json:"encrypted,omitempty"`
	Externalauth        string `json:"externalauth,omitempty"`
	Password            string `json:"password,omitempty"`
	Promptinheritedfrom string `json:"promptinheritedfrom,omitempty"`
	Promptstring        string `json:"promptstring,omitempty"`
	Timeout             int    `json:"timeout,omitempty"`
	Timeoutkind         string `json:"timeoutkind,omitempty"`
	Username            string `json:"username,omitempty"`
}
