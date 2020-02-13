package ns

type Nsweblogparam struct {
	Buffersizemb  int         `json:"buffersizemb,omitempty"`
	Builtin       interface{} `json:"builtin,omitempty"`
	Customreqhdrs interface{} `json:"customreqhdrs,omitempty"`
	Customrsphdrs interface{} `json:"customrsphdrs,omitempty"`
	Feature       string      `json:"feature,omitempty"`
}
