package snmp

type Snmpoption struct {
	Partitionnameintrap  string `json:"partitionnameintrap,omitempty"`
	Snmpset              string `json:"snmpset,omitempty"`
	Snmptraplogging      string `json:"snmptraplogging,omitempty"`
	Snmptraplogginglevel string `json:"snmptraplogginglevel,omitempty"`
}
