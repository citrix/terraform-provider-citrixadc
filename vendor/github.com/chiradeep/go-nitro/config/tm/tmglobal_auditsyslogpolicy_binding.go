package tm

type Tmglobalauditsyslogpolicybinding struct {
	Bindpolicytype int    `json:"bindpolicytype,omitempty"`
	Policyname     string `json:"policyname,omitempty"`
	Priority       int    `json:"priority,omitempty"`
}
