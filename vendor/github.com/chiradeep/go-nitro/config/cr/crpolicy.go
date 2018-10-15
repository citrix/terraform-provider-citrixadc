package cr

type Crpolicy struct {
	Action       string      `json:"action,omitempty"`
	Activepolicy bool        `json:"activepolicy,omitempty"`
	Bindhits     int         `json:"bindhits,omitempty"`
	Builtin      interface{} `json:"builtin,omitempty"`
	Cspolicytype string      `json:"cspolicytype,omitempty"`
	Domain       string      `json:"domain,omitempty"`
	Hits         int         `json:"hits,omitempty"`
	Isdefault    bool        `json:"isdefault,omitempty"`
	Labelname    string      `json:"labelname,omitempty"`
	Labeltype    string      `json:"labeltype,omitempty"`
	Logaction    string      `json:"logaction,omitempty"`
	Newname      string      `json:"newname,omitempty"`
	Policyname   string      `json:"policyname,omitempty"`
	Priority     int         `json:"priority,omitempty"`
	Rule         string      `json:"rule,omitempty"`
	Vstype       int         `json:"vstype,omitempty"`
}
