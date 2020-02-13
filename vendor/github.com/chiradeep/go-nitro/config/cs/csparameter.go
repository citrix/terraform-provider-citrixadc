package cs

type Csparameter struct {
	Builtin     interface{} `json:"builtin,omitempty"`
	Feature     string      `json:"feature,omitempty"`
	Stateupdate string      `json:"stateupdate,omitempty"`
}
