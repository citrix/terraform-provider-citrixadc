package authentication

type Authenticationpolicy struct {
	Action        string `json:"action,omitempty"`
	Comment       string `json:"comment,omitempty"`
	Description   string `json:"description,omitempty"`
	Hits          int    `json:"hits,omitempty"`
	Logaction     string `json:"logaction,omitempty"`
	Name          string `json:"name,omitempty"`
	Newname       string `json:"newname,omitempty"`
	Policysubtype string `json:"policysubtype,omitempty"`
	Rule          string `json:"rule,omitempty"`
	Undefaction   string `json:"undefaction,omitempty"`
}
