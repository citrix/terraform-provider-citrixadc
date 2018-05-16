package cache

type Cachepolicy struct {
	Action       string      `json:"action,omitempty"`
	Builtin      interface{} `json:"builtin,omitempty"`
	Flags        int         `json:"flags,omitempty"`
	Hits         int         `json:"hits,omitempty"`
	Invalgroups  interface{} `json:"invalgroups,omitempty"`
	Invalobjects interface{} `json:"invalobjects,omitempty"`
	Newname      string      `json:"newname,omitempty"`
	Policyname   string      `json:"policyname,omitempty"`
	Rule         string      `json:"rule,omitempty"`
	Storeingroup string      `json:"storeingroup,omitempty"`
	Undefaction  string      `json:"undefaction,omitempty"`
	Undefhits    int         `json:"undefhits,omitempty"`
}
