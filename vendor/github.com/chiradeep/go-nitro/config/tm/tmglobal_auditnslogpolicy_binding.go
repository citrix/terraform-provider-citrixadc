package tm

type Tmglobalauditnslogpolicybinding struct {
	Bindpolicytype int    `json:"bindpolicytype,omitempty"`
	Policyname     string `json:"policyname,omitempty"`
	Priority       int    `json:"priority,omitempty"`
}
