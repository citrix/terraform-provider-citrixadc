package smpp

type Smppparam struct {
	Addrnpi      int    `json:"addrnpi,omitempty"`
	Addrrange    string `json:"addrrange,omitempty"`
	Addrton      int    `json:"addrton,omitempty"`
	Clientmode   string `json:"clientmode,omitempty"`
	Msgqueue     string `json:"msgqueue,omitempty"`
	Msgqueuesize int    `json:"msgqueuesize,omitempty"`
}
