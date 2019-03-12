package ca

type Caaction struct {
	Accumressize int         `json:"accumressize,omitempty"`
	Builtin      interface{} `json:"builtin,omitempty"`
	Comment      string      `json:"comment,omitempty"`
	Hits         int         `json:"hits,omitempty"`
	Isdefault    bool        `json:"isdefault,omitempty"`
	Lbvserver    string      `json:"lbvserver,omitempty"`
	Name         string      `json:"name,omitempty"`
	Newname      string      `json:"newname,omitempty"`
	Type         string      `json:"type,omitempty"`
}
