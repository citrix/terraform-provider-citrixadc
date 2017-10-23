package snmp

type Snmpview struct {
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`
	Storagetype string `json:"storagetype,omitempty"`
	Subtree     string `json:"subtree,omitempty"`
	Type        string `json:"type,omitempty"`
}
