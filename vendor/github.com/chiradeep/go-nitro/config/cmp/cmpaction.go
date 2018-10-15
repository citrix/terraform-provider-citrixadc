package cmp

type Cmpaction struct {
	Addvaryheader   string      `json:"addvaryheader,omitempty"`
	Builtin         interface{} `json:"builtin,omitempty"`
	Cmptype         string      `json:"cmptype,omitempty"`
	Deltatype       string      `json:"deltatype,omitempty"`
	Isdefault       bool        `json:"isdefault,omitempty"`
	Name            string      `json:"name,omitempty"`
	Newname         string      `json:"newname,omitempty"`
	Varyheadervalue string      `json:"varyheadervalue,omitempty"`
}
