package lsn

type Lsnparameter struct {
	Maxmemlimit          int    `json:"maxmemlimit,omitempty"`
	Memlimit             int    `json:"memlimit,omitempty"`
	Memlimitactive       int    `json:"memlimitactive,omitempty"`
	Sessionsync          string `json:"sessionsync,omitempty"`
	Subscrsessionremoval string `json:"subscrsessionremoval,omitempty"`
}
