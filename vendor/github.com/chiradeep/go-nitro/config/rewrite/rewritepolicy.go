package rewrite

type Rewritepolicy struct {
	Action      string      `json:"action,omitempty"`
	Builtin     interface{} `json:"builtin,omitempty"`
	Comment     string      `json:"comment,omitempty"`
	Description string      `json:"description,omitempty"`
	Feature     string      `json:"feature,omitempty"`
	Hits        int         `json:"hits,omitempty"`
	Isdefault   bool        `json:"isdefault,omitempty"`
	Logaction   string      `json:"logaction,omitempty"`
	Name        string      `json:"name,omitempty"`
	Newname     string      `json:"newname,omitempty"`
	Rule        string      `json:"rule,omitempty"`
	Undefaction string      `json:"undefaction,omitempty"`
	Undefhits   int         `json:"undefhits,omitempty"`
}
