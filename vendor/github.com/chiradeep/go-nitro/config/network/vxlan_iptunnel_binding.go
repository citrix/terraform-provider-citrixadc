package network

type Vxlaniptunnelbinding struct {
	Id     int    `json:"id,omitempty"`
	Tunnel string `json:"tunnel,omitempty"`
}
