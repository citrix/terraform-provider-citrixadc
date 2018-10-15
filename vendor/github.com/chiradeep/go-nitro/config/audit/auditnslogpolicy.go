package audit

type Auditnslogpolicy struct {
	Action         string      `json:"action,omitempty"`
	Builtin        interface{} `json:"builtin,omitempty"`
	Expressiontype string      `json:"expressiontype,omitempty"`
	Name           string      `json:"name,omitempty"`
	Rule           string      `json:"rule,omitempty"`
}
