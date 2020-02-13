package rewrite

type Rewritepolicylabel struct {
	Builtin                interface{} `json:"builtin,omitempty"`
	Comment                string      `json:"comment,omitempty"`
	Description            string      `json:"description,omitempty"`
	Feature                string      `json:"feature,omitempty"`
	Flowtype               int         `json:"flowtype,omitempty"`
	Gotopriorityexpression string      `json:"gotopriorityexpression,omitempty"`
	Hits                   int         `json:"hits,omitempty"`
	Invokelabelname        string      `json:"invoke_labelname,omitempty"`
	Isdefault              bool        `json:"isdefault,omitempty"`
	Labelname              string      `json:"labelname,omitempty"`
	Labeltype              string      `json:"labeltype,omitempty"`
	Newname                string      `json:"newname,omitempty"`
	Numpol                 int         `json:"numpol,omitempty"`
	Priority               int         `json:"priority,omitempty"`
	Transform              string      `json:"transform,omitempty"`
}
