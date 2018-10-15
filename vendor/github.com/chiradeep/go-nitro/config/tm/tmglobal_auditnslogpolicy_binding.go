package tm

type Tmglobalauditnslogpolicybinding struct {
	Bindpolicytype         int    `json:"bindpolicytype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
