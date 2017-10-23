package appfw

type Appfwglobalauditsyslogpolicybinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Policytype             string `json:"policytype,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	State                  string `json:"state,omitempty"`
	Type                   string `json:"type,omitempty"`
}
