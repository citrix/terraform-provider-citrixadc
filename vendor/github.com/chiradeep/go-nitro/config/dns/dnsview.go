package dns

type Dnsview struct {
	Flags    int    `json:"flags,omitempty"`
	Viewname string `json:"viewname,omitempty"`
}
