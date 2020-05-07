package utility

type Ping struct {
	C        int    `json:"c,omitempty"`
	HostName string `json:"hostName,omitempty"`
	I        int    `json:"i,omitempty"`
	N        bool   `json:"n,omitempty"`
	P        string `json:"p,omitempty"`
	Q        bool   `json:"q,omitempty"`
	Response string `json:"response,omitempty"`
	S        int    `json:"s,omitempty"`
	T        int    `json:"T,omitempty"`
}
