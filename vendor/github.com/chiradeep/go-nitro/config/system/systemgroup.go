package system

type Systemgroup struct {
	Groupname    string `json:"groupname,omitempty"`
	Promptstring string `json:"promptstring,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
}
