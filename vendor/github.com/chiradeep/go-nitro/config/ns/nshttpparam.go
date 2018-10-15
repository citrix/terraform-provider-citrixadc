package ns

type Nshttpparam struct {
	Builtin          interface{} `json:"builtin,omitempty"`
	Conmultiplex     string      `json:"conmultiplex,omitempty"`
	Dropinvalreqs    string      `json:"dropinvalreqs,omitempty"`
	Http2serverside  string      `json:"http2serverside,omitempty"`
	Insnssrvrhdr     string      `json:"insnssrvrhdr,omitempty"`
	Logerrresp       string      `json:"logerrresp,omitempty"`
	Markconnreqinval string      `json:"markconnreqinval,omitempty"`
	Markhttp09inval  string      `json:"markhttp09inval,omitempty"`
	Maxreusepool     int         `json:"maxreusepool,omitempty"`
	Nssrvrhdr        string      `json:"nssrvrhdr,omitempty"`
}
