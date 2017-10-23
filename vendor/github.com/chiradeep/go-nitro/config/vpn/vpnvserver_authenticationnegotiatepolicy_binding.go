package vpn

type Vpnvserverauthenticationnegotiatepolicybinding struct {
	Acttype   int    `json:"acttype,omitempty"`
	Name      string `json:"name,omitempty"`
	Policy    string `json:"policy,omitempty"`
	Priority  int    `json:"priority,omitempty"`
	Secondary bool   `json:"secondary,omitempty"`
}
