package authorization

type Authorizationpolicy struct {
	Action       string `json:"action,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
	Newname      string `json:"newname,omitempty"`
	Rule         string `json:"rule,omitempty"`
}
