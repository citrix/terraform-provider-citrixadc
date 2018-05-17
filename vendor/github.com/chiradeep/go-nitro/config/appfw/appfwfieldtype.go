package appfw

type Appfwfieldtype struct {
	Builtin  interface{} `json:"builtin,omitempty"`
	Comment  string      `json:"comment,omitempty"`
	Name     string      `json:"name,omitempty"`
	Priority int         `json:"priority,omitempty"`
	Regex    string      `json:"regex,omitempty"`
}
