package authentication

type Authenticationvserverauthenticationcertpolicybinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}
