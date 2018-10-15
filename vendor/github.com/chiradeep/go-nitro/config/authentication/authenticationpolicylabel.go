package authentication

type Authenticationpolicylabel struct {
	Comment                string `json:"comment,omitempty"`
	Description            string `json:"description,omitempty"`
	Flowtype               int    `json:"flowtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Loginschema            string `json:"loginschema,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Type                   string `json:"type,omitempty"`
}
