package ica

type Icalatencyprofile struct {
	Builtin                  interface{} `json:"builtin,omitempty"`
	Isdefault                bool        `json:"isdefault,omitempty"`
	L7latencymaxnotifycount  int         `json:"l7latencymaxnotifycount,omitempty"`
	L7latencymonitoring      string      `json:"l7latencymonitoring,omitempty"`
	L7latencynotifyinterval  int         `json:"l7latencynotifyinterval,omitempty"`
	L7latencythresholdfactor int         `json:"l7latencythresholdfactor,omitempty"`
	L7latencywaittime        int         `json:"l7latencywaittime,omitempty"`
	Name                     string      `json:"name,omitempty"`
	Refcnt                   int         `json:"refcnt,omitempty"`
}
