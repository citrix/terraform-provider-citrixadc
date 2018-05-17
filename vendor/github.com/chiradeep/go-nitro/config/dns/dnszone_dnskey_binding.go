package dns

type Dnszonednskeybinding struct {
	Expires          int         `json:"expires,omitempty"`
	Keyname          interface{} `json:"keyname,omitempty"`
	Siginceptiontime interface{} `json:"siginceptiontime,omitempty"`
	Signed           int         `json:"signed,omitempty"`
	Zonename         string      `json:"zonename,omitempty"`
}
