package authentication

type Authenticationvserverauthenticationsamlpolicybinding struct {
	Acttype         int    `json:"acttype,omitempty"`
	Groupextraction bool   `json:"groupextraction,omitempty"`
	Name            string `json:"name,omitempty"`
	Policy          string `json:"policy,omitempty"`
	Priority        int    `json:"priority,omitempty"`
	Secondary       bool   `json:"secondary,omitempty"`
}
