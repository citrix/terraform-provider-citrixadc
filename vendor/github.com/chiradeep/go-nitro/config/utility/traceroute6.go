package utility

type Traceroute6 struct {
	Host      string `json:"host,omitempty"`
	I         bool   `json:"I,omitempty"`
	M         int    `json:"m,omitempty"`
	N         bool   `json:"n,omitempty"`
	P         int    `json:"p,omitempty"`
	Packetlen int    `json:"packetlen,omitempty"`
	Q         int    `json:"q,omitempty"`
	R         bool   `json:"r,omitempty"`
	Response  string `json:"response,omitempty"`
	S         string `json:"s,omitempty"`
	T         int    `json:"T,omitempty"`
	V         bool   `json:"v,omitempty"`
	W         int    `json:"w,omitempty"`
}
