package authentication

type Authenticationpolicylabelauthenticationpolicybinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
