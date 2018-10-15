package vpn

type Vpnglobalauthenticationldappolicybinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}
