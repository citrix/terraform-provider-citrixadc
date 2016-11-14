package basic

type Extendedmemoryparam struct {
	Maxmemlimit       int `json:"maxmemlimit,omitempty"`
	Memlimit          int `json:"memlimit,omitempty"`
	Memlimitactive    int `json:"memlimitactive,omitempty"`
	Minrequiredmemory int `json:"minrequiredmemory,omitempty"`
}
