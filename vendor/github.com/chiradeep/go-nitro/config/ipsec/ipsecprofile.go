package ipsec

type Ipsecprofile struct {
	Builtin               interface{} `json:"builtin,omitempty"`
	Encalgo               interface{} `json:"encalgo,omitempty"`
	Hashalgo              interface{} `json:"hashalgo,omitempty"`
	Ikeretryinterval      int         `json:"ikeretryinterval,omitempty"`
	Ikeversion            string      `json:"ikeversion,omitempty"`
	Lifetime              int         `json:"lifetime,omitempty"`
	Livenesscheckinterval int         `json:"livenesscheckinterval,omitempty"`
	Name                  string      `json:"name,omitempty"`
	Peerpublickey         string      `json:"peerpublickey,omitempty"`
	Privatekey            string      `json:"privatekey,omitempty"`
	Psk                   string      `json:"psk,omitempty"`
	Publickey             string      `json:"publickey,omitempty"`
	Replaywindowsize      int         `json:"replaywindowsize,omitempty"`
	Retransmissiontime    int         `json:"retransmissiontime,omitempty"`
}
