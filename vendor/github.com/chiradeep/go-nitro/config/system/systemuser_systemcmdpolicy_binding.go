package system

type Systemusersystemcmdpolicybinding struct {
	Policyname string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	Username   string `json:"username,omitempty"`
}
