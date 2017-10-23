package filter

type Filteraction struct {
	Builtin     interface{} `json:"builtin,omitempty"`
	Isdefault   bool        `json:"isdefault,omitempty"`
	Name        string      `json:"name,omitempty"`
	Page        string      `json:"page,omitempty"`
	Qual        string      `json:"qual,omitempty"`
	Respcode    int         `json:"respcode,omitempty"`
	Servicename string      `json:"servicename,omitempty"`
	Value       string      `json:"value,omitempty"`
}
