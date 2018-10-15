package lsn

type Lsnclientnetwork6binding struct {
	Clientname string `json:"clientname,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Network    string `json:"network,omitempty"`
	Network6   string `json:"network6,omitempty"`
	Td         int    `json:"td,omitempty"`
}
