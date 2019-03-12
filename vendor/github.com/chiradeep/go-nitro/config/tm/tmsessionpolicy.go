package tm

type Tmsessionpolicy struct {
	Action                 string      `json:"action,omitempty"`
	Builtin                interface{} `json:"builtin,omitempty"`
	Expressiontype         string      `json:"expressiontype,omitempty"`
	Gotopriorityexpression string      `json:"gotopriorityexpression,omitempty"`
	Hits                   int         `json:"hits,omitempty"`
	Name                   string      `json:"name,omitempty"`
	Rule                   string      `json:"rule,omitempty"`
}
