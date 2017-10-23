package snmp

type Snmpuser struct {
	Authpasswd  string `json:"authpasswd,omitempty"`
	Authtype    string `json:"authtype,omitempty"`
	Engineid    string `json:"engineid,omitempty"`
	Group       string `json:"group,omitempty"`
	Name        string `json:"name,omitempty"`
	Privpasswd  string `json:"privpasswd,omitempty"`
	Privtype    string `json:"privtype,omitempty"`
	Status      string `json:"status,omitempty"`
	Storagetype string `json:"storagetype,omitempty"`
}
