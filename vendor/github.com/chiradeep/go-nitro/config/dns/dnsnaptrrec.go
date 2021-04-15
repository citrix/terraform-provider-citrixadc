package dns

type Dnsnaptrrec struct {
	Authtype    string `json:"authtype,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Ecssubnet   string `json:"ecssubnet,omitempty"`
	Flags       string `json:"flags,omitempty"`
	Nodeid      int    `json:"nodeid,omitempty"`
	Order       int    `json:"order,omitempty"`
	Preference  int    `json:"preference,omitempty"`
	Recordid    int    `json:"recordid,omitempty"`
	Regexp      string `json:"regexp,omitempty"`
	Replacement string `json:"replacement,omitempty"`
	Services    string `json:"services,omitempty"`
	Ttl         int    `json:"ttl,omitempty"`
	Type        string `json:"type,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}
