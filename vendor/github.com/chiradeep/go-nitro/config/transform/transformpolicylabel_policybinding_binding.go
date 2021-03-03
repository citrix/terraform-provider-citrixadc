package transform

type Transformpolicylabelpolicybindingbinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
