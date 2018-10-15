package ipsec

type Ipsecparameter struct {
	Encalgo               interface{} `json:"encalgo,omitempty"`
	Hashalgo              interface{} `json:"hashalgo,omitempty"`
	Ikeretryinterval      int         `json:"ikeretryinterval,omitempty"`
	Ikeversion            string      `json:"ikeversion,omitempty"`
	Lifetime              int         `json:"lifetime,omitempty"`
	Livenesscheckinterval int         `json:"livenesscheckinterval,omitempty"`
	Perfectforwardsecrecy string      `json:"perfectforwardsecrecy,omitempty"`
	Replaywindowsize      int         `json:"replaywindowsize,omitempty"`
	Responderonly         string      `json:"responderonly,omitempty"`
	Retransmissiontime    int         `json:"retransmissiontime,omitempty"`
}
