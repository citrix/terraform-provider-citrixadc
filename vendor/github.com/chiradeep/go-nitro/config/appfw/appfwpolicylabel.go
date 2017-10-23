package appfw

type Appfwpolicylabel struct {
	Description            string `json:"description,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Policylabeltype        string `json:"policylabeltype,omitempty"`
	Policytype             string `json:"policytype,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
