package ns

type Nsvariable struct {
	Comment        string `json:"comment,omitempty"`
	Expires        int    `json:"expires,omitempty"`
	Iffull         string `json:"iffull,omitempty"`
	Ifnovalue      string `json:"ifnovalue,omitempty"`
	Ifvaluetoobig  string `json:"ifvaluetoobig,omitempty"`
	Init           string `json:"init,omitempty"`
	Name           string `json:"name,omitempty"`
	Referencecount int    `json:"referencecount,omitempty"`
	Scope          string `json:"scope,omitempty"`
	Type           string `json:"type,omitempty"`
}
