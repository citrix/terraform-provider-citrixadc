package network

type Forwardingsession struct {
	Acl6name     string `json:"acl6name,omitempty"`
	Aclname      string `json:"aclname,omitempty"`
	Connfailover string `json:"connfailover,omitempty"`
	Name         string `json:"name,omitempty"`
	Netmask      string `json:"netmask,omitempty"`
	Network      string `json:"network,omitempty"`
}
