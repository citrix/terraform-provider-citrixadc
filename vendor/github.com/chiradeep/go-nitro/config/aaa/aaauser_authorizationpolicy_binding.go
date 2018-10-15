package aaa

type Aaauserauthorizationpolicybinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Type                   string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}
