package cmp

type Cmpglobalcmppolicybinding struct {
	Numpol     int    `json:"numpol,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Policytype string `json:"policytype,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	State      string `json:"state,omitempty"`
	Type       string `json:"type,omitempty"`
}
