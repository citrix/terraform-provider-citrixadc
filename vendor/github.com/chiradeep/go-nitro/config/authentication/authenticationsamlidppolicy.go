package authentication

type Authenticationsamlidppolicy struct {
	Action                 string `json:"action,omitempty"`
	Comment                string `json:"comment,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Logaction              string `json:"logaction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Rule                   string `json:"rule,omitempty"`
	Undefaction            string `json:"undefaction,omitempty"`
}
