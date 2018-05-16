package network

type Netprofilenatrulebinding struct {
	Name      string `json:"name,omitempty"`
	Natrule   string `json:"natrule,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Rewriteip string `json:"rewriteip,omitempty"`
}
