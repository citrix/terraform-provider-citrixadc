package protocol

type Protocolhttpband struct {
	Accesscount   interface{} `json:"accesscount,omitempty"`
	Accessratio   interface{} `json:"accessratio,omitempty"`
	Avgbandsize   interface{} `json:"avgbandsize,omitempty"`
	Banddata      interface{} `json:"banddata,omitempty"`
	Bandrange     int         `json:"bandrange,omitempty"`
	Reqbandsize   int         `json:"reqbandsize,omitempty"`
	Respbandsize  int         `json:"respbandsize,omitempty"`
	Totalbandsize interface{} `json:"totalbandsize,omitempty"`
	Type          string      `json:"type,omitempty"`
}
