package rewrite

type Rewriteaction struct {
	Builtin           interface{} `json:"builtin,omitempty"`
	Bypasssafetycheck string      `json:"bypasssafetycheck,omitempty"`
	Comment           string      `json:"comment,omitempty"`
	Description       string      `json:"description,omitempty"`
	Hits              int         `json:"hits,omitempty"`
	Isdefault         bool        `json:"isdefault,omitempty"`
	Name              string      `json:"name,omitempty"`
	Newname           string      `json:"newname,omitempty"`
	Pattern           string      `json:"pattern,omitempty"`
	Referencecount    int         `json:"referencecount,omitempty"`
	Refinesearch      string      `json:"refinesearch,omitempty"`
	Search            string      `json:"search,omitempty"`
	Stringbuilderexpr string      `json:"stringbuilderexpr,omitempty"`
	Target            string      `json:"target,omitempty"`
	Type              string      `json:"type,omitempty"`
	Undefhits         int         `json:"undefhits,omitempty"`
}
