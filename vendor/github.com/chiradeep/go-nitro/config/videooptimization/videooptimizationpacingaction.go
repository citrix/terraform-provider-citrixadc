package videooptimization

type Videooptimizationpacingaction struct {
	Builtin        interface{} `json:"builtin,omitempty"`
	Comment        string      `json:"comment,omitempty"`
	Hits           int         `json:"hits,omitempty"`
	Name           string      `json:"name,omitempty"`
	Newname        string      `json:"newname,omitempty"`
	Rate           int         `json:"rate,omitempty"`
	Referencecount int         `json:"referencecount,omitempty"`
	Undefhits      int         `json:"undefhits,omitempty"`
}
