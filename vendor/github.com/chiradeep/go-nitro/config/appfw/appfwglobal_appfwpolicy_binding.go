package appfw

type Appfwglobalappfwpolicybinding struct {
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Policytype             string `json:"policytype,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	State                  string `json:"state,omitempty"`
	Type                   string `json:"type,omitempty"`
}
