package vpn

type Vpnglobalauthenticationnegotiatepolicybinding struct {
	Policyname string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	Secondary  bool   `json:"secondary,omitempty"`
}
