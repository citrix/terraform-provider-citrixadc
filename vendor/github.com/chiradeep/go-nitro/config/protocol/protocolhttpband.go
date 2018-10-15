package protocol

type Protocolhttpband struct {
	Accesscount    interface{} `json:"accesscount,omitempty"`
	Accessratio    interface{} `json:"accessratio,omitempty"`
	Accessrationew interface{} `json:"accessrationew,omitempty"`
	Avgbandsize    interface{} `json:"avgbandsize,omitempty"`
	Avgbandsizenew interface{} `json:"avgbandsizenew,omitempty"`
	Banddata       interface{} `json:"banddata,omitempty"`
	Banddatanew    interface{} `json:"banddatanew,omitempty"`
	Bandrange      int         `json:"bandrange,omitempty"`
	Nodeid         int         `json:"nodeid,omitempty"`
	Numberofbands  int         `json:"numberofbands,omitempty"`
	Reqbandsize    int         `json:"reqbandsize,omitempty"`
	Respbandsize   int         `json:"respbandsize,omitempty"`
	Totalbandsize  interface{} `json:"totalbandsize,omitempty"`
	Totals         interface{} `json:"totals,omitempty"`
	Type           string      `json:"type,omitempty"`
}
