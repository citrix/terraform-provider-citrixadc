package snmp

type Snmpmanager struct {
	Domain             string `json:"domain,omitempty"`
	Domainresolveretry int    `json:"domainresolveretry,omitempty"`
	Ip                 string `json:"ip,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
}
