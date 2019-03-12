package videooptimization

type Videooptimizationdetectionaction struct {
	Builtin        interface{} `json:"builtin,omitempty"`
	Comment        string      `json:"comment,omitempty"`
	Hits           int         `json:"hits,omitempty"`
	Name           string      `json:"name,omitempty"`
	Newname        string      `json:"newname,omitempty"`
	Referencecount int         `json:"referencecount,omitempty"`
	Type           string      `json:"type,omitempty"`
	Undefhits      int         `json:"undefhits,omitempty"`
}
