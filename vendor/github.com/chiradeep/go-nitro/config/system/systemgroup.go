package system

type Systemgroup struct {
	Allowedmanagementinterface interface{} `json:"allowedmanagementinterface,omitempty"`
	Groupname                  string      `json:"groupname,omitempty"`
	Promptstring               string      `json:"promptstring,omitempty"`
	Timeout                    int         `json:"timeout,omitempty"`
}
