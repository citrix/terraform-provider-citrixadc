package responder

type Responderaction struct {
	Builtin            interface{} `json:"builtin,omitempty"`
	Bypasssafetycheck  string      `json:"bypasssafetycheck,omitempty"`
	Comment            string      `json:"comment,omitempty"`
	Hits               int         `json:"hits,omitempty"`
	Htmlpage           string      `json:"htmlpage,omitempty"`
	Name               string      `json:"name,omitempty"`
	Newname            string      `json:"newname,omitempty"`
	Reasonphrase       string      `json:"reasonphrase,omitempty"`
	Referencecount     int         `json:"referencecount,omitempty"`
	Responsestatuscode int         `json:"responsestatuscode,omitempty"`
	Target             string      `json:"target,omitempty"`
	Type               string      `json:"type,omitempty"`
	Undefhits          int         `json:"undefhits,omitempty"`
}
