package utility

type Traceroute struct {
	Host      string `json:"host,omitempty"`
	M         int    `json:"M,omitempty"`
	N         bool   `json:"n,omitempty"`
	P         int    `json:"P,omitempty"`
	Packetlen int    `json:"packetlen,omitempty"`
	Q         int    `json:"q,omitempty"`
	R         bool   `json:"r,omitempty"`
	Response  string `json:"response,omitempty"`
	S         string `json:"S,omitempty"`
	T         int    `json:"T,omitempty"`
	V         bool   `json:"v,omitempty"`
	W         int    `json:"w,omitempty"`
}
