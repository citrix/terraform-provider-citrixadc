package tm

type Tmglobaltmtrafficpolicybinding struct {
	Bindpolicytype int    `json:"bindpolicytype,omitempty"`
	Policyname     string `json:"policyname,omitempty"`
	Priority       int    `json:"priority,omitempty"`
	Type           string `json:"type,omitempty"`
}
