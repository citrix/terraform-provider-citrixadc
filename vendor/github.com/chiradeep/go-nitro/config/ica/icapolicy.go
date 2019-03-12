package ica

type Icapolicy struct {
	Action    string      `json:"action,omitempty"`
	Builtin   interface{} `json:"builtin,omitempty"`
	Comment   string      `json:"comment,omitempty"`
	Hits      int         `json:"hits,omitempty"`
	Logaction string      `json:"logaction,omitempty"`
	Name      string      `json:"name,omitempty"`
	Newname   string      `json:"newname,omitempty"`
	Rule      string      `json:"rule,omitempty"`
	Undefhits int         `json:"undefhits,omitempty"`
}
