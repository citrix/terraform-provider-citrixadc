package ns

type Nscqaparam struct {
	Harqretxdelay int     `json:"harqretxdelay,omitempty"`
	Lr1coeflist   string  `json:"lr1coeflist,omitempty"`
	Lr1probthresh float64 `json:"lr1probthresh,omitempty"`
	Lr2coeflist   string  `json:"lr2coeflist,omitempty"`
	Lr2probthresh float64 `json:"lr2probthresh,omitempty"`
	Minrttnet1    int     `json:"minrttnet1,omitempty"`
	Minrttnet2    int     `json:"minrttnet2,omitempty"`
	Minrttnet3    int     `json:"minrttnet3,omitempty"`
	Net1cclscale  string  `json:"net1cclscale,omitempty"`
	Net1csqscale  string  `json:"net1csqscale,omitempty"`
	Net1label     string  `json:"net1label,omitempty"`
	Net1logcoef   string  `json:"net1logcoef,omitempty"`
	Net2cclscale  string  `json:"net2cclscale,omitempty"`
	Net2csqscale  string  `json:"net2csqscale,omitempty"`
	Net2label     string  `json:"net2label,omitempty"`
	Net2logcoef   string  `json:"net2logcoef,omitempty"`
	Net3cclscale  string  `json:"net3cclscale,omitempty"`
	Net3csqscale  string  `json:"net3csqscale,omitempty"`
	Net3label     string  `json:"net3label,omitempty"`
	Net3logcoef   string  `json:"net3logcoef,omitempty"`
}
