package system

type Systemgroupsystemcmdpolicybinding struct {
	Groupname  string `json:"groupname,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
}
