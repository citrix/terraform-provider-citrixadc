package filter

type Filterhtmlinjectionvariable struct {
	Builtin  interface{} `json:"builtin,omitempty"`
	Type     string      `json:"type,omitempty"`
	Value    string      `json:"value,omitempty"`
	Variable string      `json:"variable,omitempty"`
}
