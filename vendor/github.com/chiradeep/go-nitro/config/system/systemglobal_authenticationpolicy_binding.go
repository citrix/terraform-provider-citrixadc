package system

type Systemglobalauthenticationpolicybinding struct {
	Builtin                interface{} `json:"builtin,omitempty"`
	Globalbindtype         string      `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string      `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string      `json:"nextfactor,omitempty"`
	Policyname             string      `json:"policyname,omitempty"`
	Priority               int         `json:"priority,omitempty"`
}
