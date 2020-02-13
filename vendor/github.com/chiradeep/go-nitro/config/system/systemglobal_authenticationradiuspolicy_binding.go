package system

type Systemglobalauthenticationradiuspolicybinding struct {
	Builtin                interface{} `json:"builtin,omitempty"`
	Feature                string      `json:"feature,omitempty"`
	Globalbindtype         string      `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string      `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string      `json:"nextfactor,omitempty"`
	Policyname             string      `json:"policyname,omitempty"`
	Priority               int         `json:"priority,omitempty"`
}
