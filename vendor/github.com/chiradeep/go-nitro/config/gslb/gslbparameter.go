package gslb

type Gslbparameter struct {
	Dropldnsreq      string      `json:"dropldnsreq,omitempty"`
	Flags            int         `json:"flags,omitempty"`
	Ldnsentrytimeout int         `json:"ldnsentrytimeout,omitempty"`
	Ldnsmask         string      `json:"ldnsmask,omitempty"`
	Ldnsprobeorder   interface{} `json:"ldnsprobeorder,omitempty"`
	Rtttolerance     int         `json:"rtttolerance,omitempty"`
	V6ldnsmasklen    int         `json:"v6ldnsmasklen,omitempty"`
}
