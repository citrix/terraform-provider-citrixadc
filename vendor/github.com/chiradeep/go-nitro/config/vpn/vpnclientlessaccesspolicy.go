package vpn

type Vpnclientlessaccesspolicy struct {
	Builtin     interface{} `json:"builtin,omitempty"`
	Description string      `json:"description,omitempty"`
	Hits        int         `json:"hits,omitempty"`
	Isdefault   bool        `json:"isdefault,omitempty"`
	Name        string      `json:"name,omitempty"`
	Profilename string      `json:"profilename,omitempty"`
	Rule        string      `json:"rule,omitempty"`
	Undefaction string      `json:"undefaction,omitempty"`
	Undefhits   int         `json:"undefhits,omitempty"`
}
