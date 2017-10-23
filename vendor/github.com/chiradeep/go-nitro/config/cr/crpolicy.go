package cr

type Crpolicy struct {
	Builtin      interface{} `json:"builtin,omitempty"`
	Cspolicytype string      `json:"cspolicytype,omitempty"`
	Domain       string      `json:"domain,omitempty"`
	Policyname   string      `json:"policyname,omitempty"`
	Rule         string      `json:"rule,omitempty"`
	Vstype       int         `json:"vstype,omitempty"`
}
