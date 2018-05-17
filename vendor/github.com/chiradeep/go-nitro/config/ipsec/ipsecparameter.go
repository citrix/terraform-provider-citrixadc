package ipsec

type Ipsecparameter struct {
	Encalgo               interface{} `json:"encalgo,omitempty"`
	Hashalgo              interface{} `json:"hashalgo,omitempty"`
	Ikeretryinterval      int         `json:"ikeretryinterval,omitempty"`
	Ikeversion            string      `json:"ikeversion,omitempty"`
	Lifetime              int         `json:"lifetime,omitempty"`
	Livenesscheckinterval int         `json:"livenesscheckinterval,omitempty"`
	Replaywindowsize      int         `json:"replaywindowsize,omitempty"`
	Retransmissiontime    int         `json:"retransmissiontime,omitempty"`
}
