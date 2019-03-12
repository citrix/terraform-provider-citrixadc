package cache

type Cacheselector struct {
	Builtin      interface{} `json:"builtin,omitempty"`
	Flags        int         `json:"flags,omitempty"`
	Rule         interface{} `json:"rule,omitempty"`
	Selectorname string      `json:"selectorname,omitempty"`
}
