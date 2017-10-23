package ssl

type Sslpolicylabel struct {
	Description            string `json:"description,omitempty"`
	Flowtype               int    `json:"flowtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Type                   string `json:"type,omitempty"`
}
