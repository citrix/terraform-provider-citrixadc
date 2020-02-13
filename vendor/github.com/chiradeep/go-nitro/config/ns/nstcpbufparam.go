package ns

type Nstcpbufparam struct {
	Builtin  interface{} `json:"builtin,omitempty"`
	Feature  string      `json:"feature,omitempty"`
	Memlimit int         `json:"memlimit,omitempty"`
	Size     int         `json:"size,omitempty"`
}
