package system

type Systemusersystemgroupbinding struct {
	Groupname  string `json:"groupname,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	Username   string `json:"username,omitempty"`
}
