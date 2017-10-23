package cr

type Craction struct {
	Builtin        interface{} `json:"builtin,omitempty"`
	Comment        string      `json:"comment,omitempty"`
	Crtype         string      `json:"crtype,omitempty"`
	Hits           int         `json:"hits,omitempty"`
	Isdefault      bool        `json:"isdefault,omitempty"`
	Name           string      `json:"name,omitempty"`
	Referencecount string      `json:"referencecount,omitempty"`
	Undefhits      int         `json:"undefhits,omitempty"`
}
