package system

type Systemsshkey struct {
	Name       string `json:"name,omitempty"`
	Src        string `json:"src,omitempty"`
	Sshkeytype string `json:"sshkeytype,omitempty"`
}
