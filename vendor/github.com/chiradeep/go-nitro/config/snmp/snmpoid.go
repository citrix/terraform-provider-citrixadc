package snmp

type Snmpoid struct {
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
	Snmpoid    string `json:"Snmpoid,omitempty"`
}
