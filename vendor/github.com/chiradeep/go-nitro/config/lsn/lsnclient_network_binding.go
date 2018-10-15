package lsn

type Lsnclientnetworkbinding struct {
	Clientname string `json:"clientname,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Network    string `json:"network,omitempty"`
	Td         int    `json:"td,omitempty"`
}
