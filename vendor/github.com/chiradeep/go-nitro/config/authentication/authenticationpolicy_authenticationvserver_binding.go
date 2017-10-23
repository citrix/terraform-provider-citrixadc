package authentication

type Authenticationpolicyauthenticationvserverbinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
