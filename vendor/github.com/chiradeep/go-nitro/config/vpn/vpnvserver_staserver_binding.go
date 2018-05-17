package vpn

type Vpnvserverstaserverbinding struct {
	Acttype   int    `json:"acttype,omitempty"`
	Name      string `json:"name,omitempty"`
	Staauthid string `json:"staauthid,omitempty"`
	Staserver string `json:"staserver,omitempty"`
}
