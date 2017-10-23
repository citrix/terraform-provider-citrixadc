package vpn

type Vpnformssoaction struct {
	Actionurl      string `json:"actionurl,omitempty"`
	Name           string `json:"name,omitempty"`
	Namevaluepair  string `json:"namevaluepair,omitempty"`
	Nvtype         string `json:"nvtype,omitempty"`
	Passwdfield    string `json:"passwdfield,omitempty"`
	Responsesize   int    `json:"responsesize,omitempty"`
	Ssosuccessrule string `json:"ssosuccessrule,omitempty"`
	Submitmethod   string `json:"submitmethod,omitempty"`
	Userfield      string `json:"userfield,omitempty"`
}
