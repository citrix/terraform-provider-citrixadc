package authorization

type Authorizationpolicy struct {
	Action         string `json:"action,omitempty"`
	Activepolicy   int    `json:"activepolicy,omitempty"`
	Expressiontype string `json:"expressiontype,omitempty"`
	Hits           int    `json:"hits,omitempty"`
	Name           string `json:"name,omitempty"`
	Newname        string `json:"newname,omitempty"`
	Rule           string `json:"rule,omitempty"`
}
