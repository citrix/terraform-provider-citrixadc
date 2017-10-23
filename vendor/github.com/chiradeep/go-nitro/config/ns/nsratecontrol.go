package ns

type Nsratecontrol struct {
	Icmpthreshold   int `json:"icmpthreshold,omitempty"`
	Tcprstthreshold int `json:"tcprstthreshold,omitempty"`
	Tcpthreshold    int `json:"tcpthreshold,omitempty"`
	Udpthreshold    int `json:"udpthreshold,omitempty"`
}
