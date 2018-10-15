package vpn

type Vpntrafficpolicy struct {
	Action         string `json:"action,omitempty"`
	Expressiontype string `json:"expressiontype,omitempty"`
	Hits           int    `json:"hits,omitempty"`
	Name           string `json:"name,omitempty"`
	Rule           string `json:"rule,omitempty"`
}
