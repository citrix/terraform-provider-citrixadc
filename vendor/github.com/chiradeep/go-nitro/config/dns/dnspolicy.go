package dns

type Dnspolicy struct {
	Actionname        string      `json:"actionname,omitempty"`
	Builtin           interface{} `json:"builtin,omitempty"`
	Cachebypass       string      `json:"cachebypass,omitempty"`
	Description       string      `json:"description,omitempty"`
	Drop              string      `json:"drop,omitempty"`
	Hits              int         `json:"hits,omitempty"`
	Logaction         string      `json:"logaction,omitempty"`
	Name              string      `json:"name,omitempty"`
	Preferredlocation string      `json:"preferredlocation,omitempty"`
	Preferredloclist  interface{} `json:"preferredloclist,omitempty"`
	Rule              string      `json:"rule,omitempty"`
	Undefhits         int         `json:"undefhits,omitempty"`
	Viewname          string      `json:"viewname,omitempty"`
}
