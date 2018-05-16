package policy

type Policyexpression struct {
	Builtin               interface{} `json:"builtin,omitempty"`
	Clientsecuritymessage string      `json:"clientsecuritymessage,omitempty"`
	Comment               string      `json:"comment,omitempty"`
	Description           string      `json:"description,omitempty"`
	Hits                  int         `json:"hits,omitempty"`
	Isdefault             bool        `json:"isdefault,omitempty"`
	Name                  string      `json:"name,omitempty"`
	Pihits                int         `json:"pihits,omitempty"`
	Type                  string      `json:"type,omitempty"`
	Type1                 string      `json:"type1,omitempty"`
	Value                 string      `json:"value,omitempty"`
}
