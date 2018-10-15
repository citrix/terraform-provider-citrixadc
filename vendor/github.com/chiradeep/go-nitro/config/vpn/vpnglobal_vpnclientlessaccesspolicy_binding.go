package vpn

type Vpnglobalvpnclientlessaccesspolicybinding struct {
	Builtin                interface{} `json:"builtin,omitempty"`
	Globalbindtype         string      `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string      `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool        `json:"groupextraction,omitempty"`
	Policyname             string      `json:"policyname,omitempty"`
	Priority               int         `json:"priority,omitempty"`
	Secondary              bool        `json:"secondary,omitempty"`
	Type                   string      `json:"type,omitempty"`
}
