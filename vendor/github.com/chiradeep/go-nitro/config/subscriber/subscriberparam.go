package subscriber

type Subscriberparam struct {
	Builtin              interface{} `json:"builtin,omitempty"`
	Idleaction           string      `json:"idleaction,omitempty"`
	Idlettl              int         `json:"idlettl,omitempty"`
	Interfacetype        string      `json:"interfacetype,omitempty"`
	Ipv6prefixlookuplist interface{} `json:"ipv6prefixlookuplist,omitempty"`
	Keytype              string      `json:"keytype,omitempty"`
}
