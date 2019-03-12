package ica

type Icaaction struct {
	Accessprofilename  string      `json:"accessprofilename,omitempty"`
	Builtin            interface{} `json:"builtin,omitempty"`
	Hits               int         `json:"hits,omitempty"`
	Isdefault          bool        `json:"isdefault,omitempty"`
	Latencyprofilename string      `json:"latencyprofilename,omitempty"`
	Name               string      `json:"name,omitempty"`
	Newname            string      `json:"newname,omitempty"`
	Referencecount     int         `json:"referencecount,omitempty"`
	Undefhits          int         `json:"undefhits,omitempty"`
}
