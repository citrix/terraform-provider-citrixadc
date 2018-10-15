package cs

type Csaction struct {
	Builtin           interface{} `json:"builtin,omitempty"`
	Comment           string      `json:"comment,omitempty"`
	Hits              int         `json:"hits,omitempty"`
	Name              string      `json:"name,omitempty"`
	Newname           string      `json:"newname,omitempty"`
	Referencecount    int         `json:"referencecount,omitempty"`
	Targetlbvserver   string      `json:"targetlbvserver,omitempty"`
	Targetvserver     string      `json:"targetvserver,omitempty"`
	Targetvserverexpr string      `json:"targetvserverexpr,omitempty"`
	Undefhits         int         `json:"undefhits,omitempty"`
}
