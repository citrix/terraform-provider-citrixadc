package dns

type Dnsaction64 struct {
	Actionname  string      `json:"actionname,omitempty"`
	Builtin     interface{} `json:"builtin,omitempty"`
	Excluderule string      `json:"excluderule,omitempty"`
	Feature     string      `json:"feature,omitempty"`
	Mappedrule  string      `json:"mappedrule,omitempty"`
	Prefix      string      `json:"prefix,omitempty"`
}
