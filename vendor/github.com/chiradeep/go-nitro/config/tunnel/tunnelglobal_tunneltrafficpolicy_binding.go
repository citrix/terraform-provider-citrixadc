package tunnel

type Tunnelglobaltunneltrafficpolicybinding struct {
	Builtin                interface{} `json:"builtin,omitempty"`
	Globalbindtype         string      `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string      `json:"gotopriorityexpression,omitempty"`
	Numpol                 int         `json:"numpol,omitempty"`
	Policyname             string      `json:"policyname,omitempty"`
	Policytype             string      `json:"policytype,omitempty"`
	Priority               int         `json:"priority,omitempty"`
	State                  string      `json:"state,omitempty"`
	Type                   string      `json:"type,omitempty"`
}
