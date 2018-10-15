package audit

type Auditsyslogglobalauditsyslogpolicybinding struct {
	Builtin        interface{} `json:"builtin,omitempty"`
	Globalbindtype string      `json:"globalbindtype,omitempty"`
	Numpol         int         `json:"numpol,omitempty"`
	Policyname     string      `json:"policyname,omitempty"`
	Priority       int         `json:"priority,omitempty"`
}
