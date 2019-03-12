package feo

type Feopolicy struct {
	Action    string      `json:"action,omitempty"`
	Builtin   interface{} `json:"builtin,omitempty"`
	Hits      int         `json:"hits,omitempty"`
	Name      string      `json:"name,omitempty"`
	Rule      string      `json:"rule,omitempty"`
	Undefhits int         `json:"undefhits,omitempty"`
}
