package cache

type Cacheselector struct {
	Flags        int         `json:"flags,omitempty"`
	Rule         interface{} `json:"rule,omitempty"`
	Selectorname string      `json:"selectorname,omitempty"`
}
