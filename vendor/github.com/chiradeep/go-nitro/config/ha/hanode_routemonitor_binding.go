package ha

type Hanoderoutemonitorbinding struct {
	Flags             int    `json:"flags,omitempty"`
	Id                int    `json:"id,omitempty"`
	Netmask           string `json:"netmask,omitempty"`
	Routemonitor      string `json:"routemonitor,omitempty"`
	Routemonitorstate string `json:"routemonitorstate,omitempty"`
}
